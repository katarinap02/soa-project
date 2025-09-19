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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TourPurchaseRepository interface {
	SaveToken(ctx context.Context, token *model.TourPurchaseToken) error
	HasUserPurchasedTour(ctx context.Context, userID string, tourID primitive.ObjectID) bool
}

type tourPurchaseRepository struct {
	coll   *mongo.Collection
	logger *log.Logger
}

// Konstruktor sa ctx, mongoURI i logger
func NewMongoTourPurchaseRepo(ctx context.Context, mongoURI string, logger *log.Logger) (TourPurchaseRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
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

	logger.Println("Connected to MongoDB for tour purchase tokens")

	coll := client.Database("soadb").Collection("tour_purchase_tokens")
	return &tourPurchaseRepository{coll: coll, logger: logger}, nil
}

func (r *tourPurchaseRepository) SaveToken(ctx context.Context, token *model.TourPurchaseToken) error {
	_, err := r.coll.InsertOne(ctx, token)
	return err
}

func (r *tourPurchaseRepository) HasUserPurchasedTour(ctx context.Context, userID string, tourID primitive.ObjectID) bool {
	count, err := r.coll.CountDocuments(ctx, bson.M{
		"userId": userID,
		"tourId": tourID,
	})
	if err != nil {
		return false
	}
	return count > 0
}
