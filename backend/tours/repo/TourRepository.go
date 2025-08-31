package repo

import (
	"context"
	"log"
	"time"

	"database-example/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type TourRepo interface {
	GetAll(ctx context.Context) ([]*model.Tour, error)
	Create(ctx context.Context, tour *model.Tour) error
}

type mongoTourRepo struct {
	collection *mongo.Collection
	logger     *log.Logger
}

func NewMongoTourRepo(ctx context.Context, uri string, logger *log.Logger) (TourRepo, error) {
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
	logger.Println("Connected to MongoDB")

	coll := client.Database("mongoDemo").Collection("tours")
	return &mongoTourRepo{collection: coll, logger: logger}, nil
}

// Metoda za testiranje konekcije i vraćanje svih tour-eva
func (m *mongoTourRepo) GetAll(ctx context.Context) ([]*model.Tour, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		m.logger.Println("Error fetching tours:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var tours []*model.Tour
	if err := cursor.All(ctx, &tours); err != nil {
		m.logger.Println("Error decoding tours:", err)
		return nil, err
	}

	return tours, nil
}

// Implementacija Create metode
func (m *mongoTourRepo) Create(ctx context.Context, tour *model.Tour) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.collection.InsertOne(ctx, tour)
	if err != nil {
		m.logger.Println("Error creating tour:", err)
		return err
	}
	return nil
}
