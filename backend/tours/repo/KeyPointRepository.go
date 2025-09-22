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

type KeyPointRepo interface {
	Create(ctx context.Context, kp *model.KeyPoint) error
	GetByTour(ctx context.Context, tourID primitive.ObjectID) ([]*model.KeyPoint, error)
	Update(ctx context.Context, id primitive.ObjectID, kp *model.KeyPoint) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type mongoKeyPointRepo struct {
	collection *mongo.Collection
	logger     *log.Logger
}

func NewMongoKeyPointRepo(ctx context.Context, uri string, logger *log.Logger) (KeyPointRepo, error) {
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
	logger.Println("Connected to MongoDB for KeyPoints")

	coll := client.Database("soadb").Collection("keypoints")
	return &mongoKeyPointRepo{collection: coll, logger: logger}, nil
}

func (m *mongoKeyPointRepo) Create(ctx context.Context, kp *model.KeyPoint) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.collection.InsertOne(ctx, kp)
	if err != nil {
		m.logger.Println("Error creating keypoint:", err)
		return err
	}
	return nil
}

func (m *mongoKeyPointRepo) GetByTour(ctx context.Context, tourID primitive.ObjectID) ([]*model.KeyPoint, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"tourId": tourID}
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		m.logger.Println("Error fetching keypoints by tour:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var kps []*model.KeyPoint
	if err := cursor.All(ctx, &kps); err != nil {
		m.logger.Println("Error decoding keypoints:", err)
		return nil, err
	}

	return kps, nil
}

func (m *mongoKeyPointRepo) Update(ctx context.Context, id primitive.ObjectID, kp *model.KeyPoint) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": kp}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err
}

func (m *mongoKeyPointRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := m.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
