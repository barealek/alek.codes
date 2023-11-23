package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Health() error
}

type service struct {
	db *mongo.Database
}

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_NAME")
)

func New() (Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))
	if err != nil {
		return nil, err
	}

	return &service{
		db: client.Database(os.Getenv(database)),
	}, nil
}

func (s *service) Health() error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	return s.db.Client().Ping(ctx, nil)
}

func (s *service) Create(ctx context.Context, collection string, data interface{}) (primitive.ObjectID, error) {
	// Create function creates a new document in the specified collection, with the supplied data.
	// Returns:
	// the ID of the newly created document (can be empty)
	// an error (can be nil)

	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	res, err := s.db.Collection(collection).InsertOne(ctx, data)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	inserted := res.InsertedID.(primitive.ObjectID)
	return inserted, nil
}
