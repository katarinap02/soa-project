package repo

import (
	"context"
	"log"
	"time"

	"database-example/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ReviewRepo interface {
	AddReview(ctx context.Context, review *model.Review) (*model.Review, error)
	GetReviewsByTour(ctx context.Context, tourId string) (model.Reviews, error)
}

type mongoReviewRepo struct {
	collection *mongo.Collection
	logger     *log.Logger
}

func NewMongoReviewRepo(ctx context.Context, uri string, logger *log.Logger) (ReviewRepo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	// Ping baze
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Ping(ctxPing, readpref.Primary()); err != nil {
		return nil, err
	}
	logger.Println("Connected to MongoDB for reviews")

	coll := client.Database("soadb").Collection("reviews")
	return &mongoReviewRepo{collection: coll, logger: logger}, nil
}

func (m *mongoReviewRepo) AddReview(ctx context.Context, review *model.Review) (*model.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	review.ID = primitive.NewObjectID()
	review.CommentDate = time.Now()

	_, err := m.collection.InsertOne(ctx, review)
	if err != nil {
		m.logger.Println("Error adding review:", err)
		return nil, err
	}
	m.logger.Println("Review added with ID:", review.ID.Hex())
	return review, nil
}

func (m *mongoReviewRepo) GetReviewsByTour(ctx context.Context, tourId string) (model.Reviews, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(tourId)
	if err != nil {
		m.logger.Println("Invalid tour ID:", tourId)
		return nil, err
	}

	cursor, err := m.collection.Find(ctx, bson.M{"tourId": oid})
	if err != nil {
		m.logger.Println("Error fetching reviews for tour:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews model.Reviews
	for cursor.Next(ctx) {
		var review model.Review
		if err := cursor.Decode(&review); err != nil {
			m.logger.Println("Error decoding review:", err)
			return nil, err
		}
		reviews = append(reviews, &review)
	}
	m.logger.Println("Fetched", len(reviews), "reviews for tour:", tourId)
	return reviews, nil
}
