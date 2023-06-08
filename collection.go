package mongoose

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=collection.go -destination=./mocks/collection.go -package=mocks
type CollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

type CollectionHelper struct {
	collection *mongo.Collection
}

func NewCollection(collectionName string, database DatabaseInterface) CollectionInterface {
	return &CollectionHelper{
		collection: database.Collection(collectionName),
	}
}

func (col CollectionHelper) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return col.collection.InsertOne(ctx, document)
}
