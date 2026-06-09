package repository

import (
	"context"

	"indo-retail-scraper/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database, collectionName string) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection(collectionName + "_products"),
	}
}

// insert batch
func (r *ProductRepository) InsertProducts(ctx context.Context, products []model.Product) error {
	var docs []interface{}

	for _, j := range products {
		docs = append(docs, j)
	}

	if len(docs) == 0 {
		return nil
	}

	_, err := r.collection.InsertMany(ctx, docs)
	return err
}
