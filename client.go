package mongoose

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=mocks
type MongoClientInterface interface {
	Ping(ctx context.Context) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
}

type clientHelper struct {
	connection *mongo.Client
}

func NewConnection(ctx context.Context, connectionString string) (MongoClientInterface, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	return &clientHelper{
		connection: conn,
	}, nil
}

func (conn *clientHelper) Ping(ctx context.Context) error {
	return conn.connection.Ping(ctx, nil)
}

func (conn *clientHelper) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return conn.connection.Database(name, opts...)
}
