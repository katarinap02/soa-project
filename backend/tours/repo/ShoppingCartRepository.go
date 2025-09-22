package repo

import (
	"context"
	"log"
	"time"

	"database-example/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ShoppingCartRepository interface {
	GetCartByUser(ctx context.Context, userId string) (*model.ShoppingCart, error)
	SaveCart(ctx context.Context, cart *model.ShoppingCart) error
	ClearCart(ctx context.Context, userId string) error
}

type shoppingCartRepository struct {
	coll   *mongo.Collection
	logger *log.Logger
}

// Konstruktor sa ctx, mongoURI i logger
func NewMongoShoppingCartRepo(ctx context.Context, mongoURI string, logger *log.Logger) (ShoppingCartRepository, error) {
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

	logger.Println("Connected to MongoDB for shopping carts")

	coll := client.Database("soadb").Collection("shopping_carts")
	return &shoppingCartRepository{coll: coll, logger: logger}, nil
}

func (r *shoppingCartRepository) GetCartByUser(ctx context.Context, userId string) (*model.ShoppingCart, error) {
	var cart model.ShoppingCart
	err := r.coll.FindOne(ctx, bson.M{"userId": userId}).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		cart = model.ShoppingCart{
			ID:     primitive.NewObjectID(),
			UserID: userId,
			Items:  []model.OrderItem{},
		}
		return &cart, nil
	}
	return &cart, err
}

func (r *shoppingCartRepository) SaveCart(ctx context.Context, cart *model.ShoppingCart) error {
	_, err := r.coll.UpdateOne(ctx,
		bson.M{"userId": cart.UserID},
		bson.M{"$set": cart},
		options.Update().SetUpsert(true),
	)
	return err
}

func (r *shoppingCartRepository) ClearCart(ctx context.Context, userId string) error {
	_, err := r.coll.DeleteOne(ctx, bson.M{"userId": userId})
	return err
}
