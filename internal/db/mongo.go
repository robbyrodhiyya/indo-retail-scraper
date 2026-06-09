package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect() (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(os.Getenv("DB_NAME"))

	return &Mongo{
		Client: client,
		DB:     db,
	}, nil
}

// dynamic collection
func (m *Mongo) GetCollection(name string) *mongo.Collection {
	return m.DB.Collection(name)
}

// setup index
func EnsureProductIndexes(ctx context.Context, col *mongo.Collection) error {
	_, err := col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{
			"source":     1,
			"product_id": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	return err
}
