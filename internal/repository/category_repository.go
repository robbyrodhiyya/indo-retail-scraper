package repository

import (
	"context"

	"indo-retail-scraper/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database, collectionName string) *CategoryRepository {
	return &CategoryRepository{
		collection: db.Collection(collectionName + "_categories"),
	}
}

func (r *CategoryRepository) InsertCategories(
	ctx context.Context,
	categories []model.Category,
) error {

	var docs []interface{}

	for _, c := range categories {
		docs = append(docs, c)
	}

	if len(docs) == 0 {
		return nil
	}

	_, err := r.collection.InsertMany(ctx, docs)

	return err
}
