package mongoose

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=database.go -destination=./mocks/database.go -package=mocks
type DatabaseInterface interface {
	Client() *mongo.Client
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
}

type databaseHelper struct {
	database *mongo.Database
}

func NewDatabase(databaseName string, client MongoClientInterface) DatabaseInterface {
	return &databaseHelper{
		database: client.Database(databaseName),
	}
}

func (db databaseHelper) Client() *mongo.Client {
	return db.database.Client()
}

func (db databaseHelper) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return db.database.Collection(name, opts...)
}

func (db databaseHelper) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return db.database.Client().UseSession(ctx, fn)
}
