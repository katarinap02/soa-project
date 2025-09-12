package repo

import (
	"context"
	"database-example/dto"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FollowerRepository struct {
	driver neo4j.DriverWithContext
	logger *log.Logger
}

func New(logger *log.Logger) (*FollowerRepository, error) {
	// Local instance
	uri := os.Getenv("NEO4J_DB")
	user := os.Getenv("NEO4J_USERNAME")
	pass := os.Getenv("NEO4J_PASS")
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	return &FollowerRepository{
		driver: driver,
		logger: logger,
	}, nil
}

// Check if connection is established
func (repo *FollowerRepository) CheckConnection() {
	ctx := context.Background()
	err := repo.driver.VerifyConnectivity(ctx)
	if err != nil {
		repo.logger.Panic(err)
		return
	}
	repo.logger.Printf(`Neo4J server address: %s`, repo.driver.Target().Host)
}

// Disconnect from database
func (repo *FollowerRepository) CloseDriverConnection(ctx context.Context) {
	repo.driver.Close(ctx)
}

// kreira FOLLOWS vezu između dva korisnika
func (repo *FollowerRepository) FollowUser(followerID, followeeID uuid.UUID) error {
	ctx := context.Background()
	session := repo.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	savedRelation, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			// Prvo proverava da li već postoji veza
			checkResult, err := transaction.Run(ctx,
				`MATCH (follower:User {id: $followerId})-[r:FOLLOWS]->(followee:User {id: $followeeId})
				 RETURN r`,
				map[string]any{
					"followerId": followerID.String(),
					"followeeId": followeeID.String(),
				})
			if err != nil {
				return nil, err
			}

			if checkResult.Next(ctx) {
				return nil, fmt.Errorf("user already follows this person")
			}

			// Kreira čvorove ako ne postoje i vezu
			result, err := transaction.Run(ctx,
				`MERGE (follower:User {id: $followerId})
				 MERGE (followee:User {id: $followeeId})
				 CREATE (follower)-[:FOLLOWS]->(followee)
				 RETURN 'Follow relationship created'`,
				map[string]any{
					"followerId": followerID.String(),
					"followeeId": followeeID.String(),
				})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})

	if err != nil {
		repo.logger.Printf("Error creating follow relationship: %v", err)
		return err
	}

	repo.logger.Printf("Follow relationship created: %v", savedRelation)
	return nil
}

// UnfollowUser - briše FOLLOWS vezu između dva korisnika
func (repo *FollowerRepository) UnfollowUser(followerID, followeeID uuid.UUID) error {
	ctx := context.Background()
	session := repo.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	deletedRelation, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`MATCH (follower:User {id: $followerId})-[r:FOLLOWS]->(followee:User {id: $followeeId})
				 DELETE r
				 RETURN 'Follow relationship deleted'`,
				map[string]any{
					"followerId": followerID.String(),
					"followeeId": followeeID.String(),
				})
			if err != nil {
				return nil, err
			}

			summary, err := result.Consume(ctx)
			if err != nil {
				return nil, err
			}

			if summary.Counters().RelationshipsDeleted() == 0 {
				return nil, fmt.Errorf("follow relationship not found")
			}

			return "Follow relationship deleted", nil
		})

	if err != nil {
		repo.logger.Printf("Error deleting follow relationship: %v", err)
		return err
	}

	repo.logger.Printf("Unfollow successful: %v", deletedRelation)
	return nil
}

// IsFollowing - proverava da li korisnik prati drugog korisnika
func (repo *FollowerRepository) IsFollowing(followerID, followeeID uuid.UUID) (bool, error) {
	ctx := context.Background()
	session := repo.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`MATCH (follower:User {id: $followerId})-[:FOLLOWS]->(followee:User {id: $followeeId})
				 RETURN true as follows`,
				map[string]any{
					"followerId": followerID.String(),
					"followeeId": followeeID.String(),
				})
			if err != nil {
				return false, err
			}

			return result.Next(ctx), nil
		})

	if err != nil {
		repo.logger.Printf("Error checking follow relationship: %v", err)
		return false, err
	}

	return result.(bool), nil
}

// GetFollowing - vraća listu korisnika koje dati korisnik prati
func (repo *FollowerRepository) GetFollowing(userID uuid.UUID) ([]uuid.UUID, error) {
	ctx := context.Background()
	session := repo.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	followingResults, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`MATCH (user:User {id: $userId})-[:FOLLOWS]->(following:User)
				 RETURN following.id as followingId`,
				map[string]any{"userId": userID.String()})
			if err != nil {
				return nil, err
			}

			var following []uuid.UUID
			for result.Next(ctx) {
				record := result.Record()
				if followingIdValue, ok := record.Get("followingId"); ok {
					if followingIdStr, ok := followingIdValue.(string); ok {
						if followingId, err := uuid.Parse(followingIdStr); err == nil {
							following = append(following, followingId)
						}
					}
				}
			}
			return following, result.Err()
		})

	if err != nil {
		repo.logger.Printf("Error querying following: %v", err)
		return nil, err
	}
	return followingResults.([]uuid.UUID), nil
}

// GetFollowers - vraća listu korisnika koji prate datog korisnika
func (repo *FollowerRepository) GetFollowers(userID uuid.UUID) ([]uuid.UUID, error) {
	ctx := context.Background()
	session := repo.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	followersResults, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`MATCH (follower:User)-[:FOLLOWS]->(user:User {id: $userId})
				 RETURN follower.id as followerId`,
				map[string]any{"userId": userID.String()})
			if err != nil {
				return nil, err
			}

			var followers []uuid.UUID
			for result.Next(ctx) {
				record := result.Record()
				if followerIdValue, ok := record.Get("followerId"); ok {
					if followerIdStr, ok := followerIdValue.(string); ok {
						if followerId, err := uuid.Parse(followerIdStr); err == nil {
							followers = append(followers, followerId)
						}
					}
				}
			}
			return followers, result.Err()
		})

	if err != nil {
		repo.logger.Printf("Error querying followers: %v", err)
		return nil, err
	}
	return followersResults.([]uuid.UUID), nil
}

// vraća preporuke koga da prati na osnovu zajedničkih veza
func (repo *FollowerRepository) GetRecommendations(userID uuid.UUID, limit int) ([]dto.RecommendationResponse, error) {
	ctx := context.Background()
	session := repo.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	recommendationResults, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`MATCH (user:User {id: $userId})-[:FOLLOWS]->(following:User)-[:FOLLOWS]->(recommended:User)
				 WHERE NOT (user)-[:FOLLOWS]->(recommended) AND recommended.id <> $userId
				 WITH recommended.id as recommendedId, COUNT(*) as score
				 ORDER BY score DESC
				 LIMIT $limit
				 RETURN recommendedId, score`,
				map[string]any{
					"userId": userID.String(),
					"limit":  limit,
				})
			if err != nil {
				return nil, err
			}

			var recommendations []dto.RecommendationResponse
			for result.Next(ctx) {
				record := result.Record()

				recommendedIdValue, _ := record.Get("recommendedId")
				scoreValue, _ := record.Get("score")

				if recommendedIdStr, ok := recommendedIdValue.(string); ok {
					if recommendedId, err := uuid.Parse(recommendedIdStr); err == nil {
						score := int(scoreValue.(int64))

						recommendation := dto.RecommendationResponse{
							UserID: recommendedId,
							Score:  score,
						}
						recommendations = append(recommendations, recommendation)
					}
				}
			}
			return recommendations, result.Err()
		})

	if err != nil {
		repo.logger.Printf("Error querying recommendations: %v", err)
		return nil, err
	}
	return recommendationResults.([]dto.RecommendationResponse), nil
}
