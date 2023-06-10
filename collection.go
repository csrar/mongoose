package mongoose

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=collection.go -destination=./mocks/collection.go -package=mocks
type CollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
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

func (col CollectionHelper) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return col.collection.UpdateOne(ctx, filter, update, opts...)
}

func (col CollectionHelper) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	return col.collection.FindOne(ctx, filter, opts...)
}
