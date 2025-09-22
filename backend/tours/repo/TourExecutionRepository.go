package repo

import (
	"context"
	"database-example/model"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type TourExecutionRepo interface {
	Create(ctx context.Context, te *model.TourExecution) error
	GetByID(ctx context.Context, id string) (*model.TourExecution, error)
	GetByTourAndTouristId(ctx context.Context, tourID, touristID string) ([]*model.TourExecution, error)
	GetActiveByTouristId(ctx context.Context, touristID string) ([]*model.TourExecution, error)
	GetByTouristId(ctx context.Context, touristID string) ([]*model.TourExecution, error)
	UpdateLastActivity(ctx context.Context, id string) error
	CompleteExecution(ctx context.Context, id string) error
	AbandonExecution(ctx context.Context, id string) error

	CreateCompletedKeyPoint(ctx context.Context, ckp *model.CompletedKeyPoint) error
	GetCompletedKeyPointsByExecution(ctx context.Context, tourExecutionID string) ([]*model.CompletedKeyPoint, error)
}

type mongoTourExecutionRepo struct {
	executionCollection         *mongo.Collection
	completedKeyPointCollection *mongo.Collection
	logger                      *log.Logger
}

func NewMongoTourExecutionRepo(ctx context.Context, uri string, logger *log.Logger) (TourExecutionRepo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Ping(ctxPing, readpref.Primary()); err != nil {
		return nil, err
	}
	logger.Println("Connected to MongoDB for TourExecutions")

	db := client.Database("soadb")
	executionColl := db.Collection("tour_executions")
	completedKeyPointColl := db.Collection("completed_key_points")

	return &mongoTourExecutionRepo{executionCollection: executionColl, completedKeyPointCollection: completedKeyPointColl, logger: logger}, nil
}

func (m *mongoTourExecutionRepo) Create(ctx context.Context, te *model.TourExecution) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Validacija da tourID i touristID nisu prazni
	if te.TourID == "" || te.TouristID == "" {
		return errors.New("tourID and touristID are required")
	}

	result, err := m.executionCollection.InsertOne(ctx, te)
	if err != nil {
		m.logger.Println("Error creating tour execution:", err)
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		te.ID = oid
	}

	return nil
}

func (m *mongoTourExecutionRepo) GetByID(ctx context.Context, id string) (*model.TourExecution, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var te model.TourExecution
	err = m.executionCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&te)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("tour execution not found")
		}
		m.logger.Println("Error fetching tour execution:", err)
		return nil, err
	}

	return &te, nil
}

func (m *mongoTourExecutionRepo) UpdateLastActivity(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"lastActivity": time.Now().UTC()}}
	_, err = m.executionCollection.UpdateOne(ctx, filter, update)
	return err
}

func (m *mongoTourExecutionRepo) CompleteExecution(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Prvo proveri da li je execution aktivan
	te, err := m.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if te.Status != model.StatusActive {
		return errors.New("tour execution is not active")
	}

	now := time.Now().UTC()
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"status":       model.StatusCompleted,
		"lastActivity": now,
	}}
	_, err = m.executionCollection.UpdateOne(ctx, filter, update)
	return err
}

func (m *mongoTourExecutionRepo) AbandonExecution(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Prvo proveri da li je execution aktivan
	te, err := m.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if te.Status != model.StatusActive {
		return errors.New("tour execution is not active")
	}

	now := time.Now().UTC()
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"status":       model.StatusAbandoned,
		"lastActivity": now,
	}}
	_, err = m.executionCollection.UpdateOne(ctx, filter, update)
	return err
}

func (m *mongoTourExecutionRepo) CreateCompletedKeyPoint(ctx context.Context, ckp *model.CompletedKeyPoint) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if ckp.TourExecutionID == "" {
		return errors.New("tour execution ID is required")
	}

	te, err := m.GetByID(ctx, ckp.TourExecutionID)
	if err != nil {
		return errors.New("tour execution not found")
	}

	if te.Status != model.StatusActive {
		return errors.New("tour execution is not active")
	}

	result, err := m.completedKeyPointCollection.InsertOne(ctx, ckp)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("key point already completed")
		}
		m.logger.Println("Error creating completed key point:", err)
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		ckp.ID = oid
	}

	if err := m.UpdateLastActivity(ctx, ckp.TourExecutionID); err != nil {
		m.logger.Println("Warning: Failed to update last activity:", err)
	}

	return nil
}

func (m *mongoTourExecutionRepo) GetCompletedKeyPointsByExecution(ctx context.Context, tourExecutionID string) ([]*model.CompletedKeyPoint, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// tourExecutionID se čuva kao string u completed_key_points kolekciji
	filter := bson.M{"tourExecutionId": tourExecutionID}
	cursor, err := m.completedKeyPointCollection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "completedTime", Value: 1}}))
	if err != nil {
		m.logger.Println("Error fetching completed key points:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var ckps []*model.CompletedKeyPoint
	if err := cursor.All(ctx, &ckps); err != nil {
		m.logger.Println("Error decoding completed key points:", err)
		return nil, err
	}

	return ckps, nil
}

func (m *mongoTourExecutionRepo) GetByTourAndTouristId(ctx context.Context, tourID, touristID string) ([]*model.TourExecution, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Validacija da parametri nisu prazni
	if tourID == "" || touristID == "" {
		return nil, errors.New("tourID and touristID are required")
	}

	filter := bson.M{
		"tourId":    tourID,    // String UUID
		"touristId": touristID, // String UUID
	}

	cursor, err := m.executionCollection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "lastActivity", Value: -1}}))
	if err != nil {
		m.logger.Println("Error fetching tour executions by tour and tourist:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var executions []*model.TourExecution
	if err := cursor.All(ctx, &executions); err != nil {
		m.logger.Println("Error decoding tour executions:", err)
		return nil, err
	}

	return executions, nil
}

func (m *mongoTourExecutionRepo) GetActiveByTouristId(ctx context.Context, touristID string) ([]*model.TourExecution, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{
		"touristId": touristID, // String UUID
		"status":    model.StatusActive,
	}

	cursor, err := m.executionCollection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "lastActivity", Value: -1}}))
	if err != nil {
		m.logger.Println("Error fetching active tour executions by tourist:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var executions []*model.TourExecution
	if err := cursor.All(ctx, &executions); err != nil {
		m.logger.Println("Error decoding active tour executions:", err)
		return nil, err
	}

	return executions, nil
}

func (m *mongoTourExecutionRepo) GetByTouristId(ctx context.Context, touristID string) ([]*model.TourExecution, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{
		"touristId": touristID, // String UUID
	}

	cursor, err := m.executionCollection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "lastActivity", Value: -1}}))
	if err != nil {
		m.logger.Println("Error fetching active tour executions by tourist:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var executions []*model.TourExecution
	if err := cursor.All(ctx, &executions); err != nil {
		m.logger.Println("Error decoding tour executions:", err)
		return nil, err
	}

	return executions, nil
}
