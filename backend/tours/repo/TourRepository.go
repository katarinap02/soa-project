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

// TourRepo interfejs
type TourRepo interface {
	GetAll(ctx context.Context) ([]*model.Tour, error)
	Create(ctx context.Context, tour *model.Tour) error
	GetByAuthor(ctx context.Context, authorID string) ([]*model.Tour, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*model.Tour, error)
}

// Mongo implementacija
type mongoTourRepo struct {
	collection *mongo.Collection
	logger     *log.Logger
}

// Kreiranje novog Mongo repo
func NewMongoTourRepo(ctx context.Context, uri string, logger *log.Logger) (TourRepo, error) {
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
	logger.Println("Connected to MongoDB")

	coll := client.Database("soadb").Collection("tours")
	return &mongoTourRepo{collection: coll, logger: logger}, nil
}

// Vraća sve ture
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

// Kreira novu turu
func (m *mongoTourRepo) Create(ctx context.Context, tour *model.Tour) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if tour.ID.IsZero() {
		tour.ID = primitive.NewObjectID()
	}

	_, err := m.collection.InsertOne(ctx, tour)
	if err != nil {
		m.logger.Println("Error creating tour:", err)
		return err
	}
	return nil
}

// Vraća ture po UUID authorID
func (m *mongoTourRepo) GetByAuthor(ctx context.Context, authorID string) ([]*model.Tour, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"authorId": authorID} // UUID kao string
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		m.logger.Println("Error fetching tours by author:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var tours []*model.Tour
	if err := cursor.All(ctx, &tours); err != nil {
		m.logger.Println("Error decoding tours by author:", err)
		return nil, err
	}
	return tours, nil
}

func (r *mongoTourRepo) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Tour, error) {
	var tour model.Tour
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&tour)
	if err != nil {
		return nil, err
	}
	return &tour, nil
}
