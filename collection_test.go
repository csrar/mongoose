package mongoose

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_NewCollection(t *testing.T) {
	conn, _ := NewConnection(context.Background(), "mongodb://mock-string")
	db := NewDatabase("mock-db", conn)
	got := NewCollection("mock-collection", db)
	assert.NotEmpty(t, got)
}

func Test_CollectionHelper_InsertOne(t *testing.T) {
	type inParams struct {
		ctxTimeout time.Duration
		document   interface{}
	}
	tests := []struct {
		name           string
		expectedErr    bool
		mongoHost      string
		collectionName string
		databaseName   string
		inParams       inParams
	}{
		{
			name:           "Test insert one with error due timeout connection",
			expectedErr:    true,
			mongoHost:      "mongodb://mock-string",
			collectionName: "mock-collection",
			databaseName:   "mock-database",
			inParams: inParams{
				ctxTimeout: time.Second * 1,
				document:   bson.M{"mock": "document"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tt.inParams.ctxTimeout)
			defer cancel()
			conn, _ := NewConnection(ctx, tt.mongoHost)
			db := NewDatabase(tt.databaseName, conn)
			collection := NewCollection(tt.collectionName, db)
			got, err := collection.InsertOne(ctx, tt.inParams.document)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_CollectionHelper_UpdateOne(t *testing.T) {
	type inParams struct {
		ctxTimeout time.Duration
		document   interface{}
		documentID string
	}
	tests := []struct {
		name           string
		expectedErr    bool
		mongoHost      string
		collectionName string
		databaseName   string
		inParams       inParams
	}{
		{
			name:           "Test update one with error due timeout connection",
			expectedErr:    true,
			mongoHost:      "mongodb://mock-string",
			collectionName: "mock-collection",
			databaseName:   "mock-database",
			inParams: inParams{
				ctxTimeout: time.Second * 1,
				document:   bson.M{"$set": bson.M{"mock": "document"}},
				documentID: "111111111111111111111111",
			},
		},
	}
	for _, tt := range tests {
		ctx, cancel := context.WithTimeout(context.Background(), tt.inParams.ctxTimeout)
		defer cancel()
		conn, _ := NewConnection(ctx, tt.mongoHost)
		db := NewDatabase(tt.databaseName, conn)
		collection := NewCollection(tt.collectionName, db)
		docID, err := primitive.ObjectIDFromHex(tt.inParams.documentID)
		assert.Nil(t, err)
		got, err := collection.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": docID}}, tt.inParams.document)
		if tt.expectedErr {
			assert.Error(t, err)
		} else {
			assert.NotNil(t, got)
		}
	}
}

func Test_CollectionHelper_FindOne(t *testing.T) {
	type inParams struct {
		ctxTimeout time.Duration
		document   interface{}
		documentID string
	}
	tests := []struct {
		name           string
		expectedErr    bool
		mongoHost      string
		collectionName string
		databaseName   string
		inParams       inParams
	}{
		{
			name:           "Test find one with error due timeout connection",
			expectedErr:    true,
			mongoHost:      "mongodb://mock-string",
			collectionName: "mock-collection",
			databaseName:   "mock-database",
			inParams: inParams{
				ctxTimeout: time.Second * 1,
				document:   bson.M{"$set": bson.M{"mock": "document"}},
				documentID: "111111111111111111111111",
			},
		},
	}
	for _, tt := range tests {
		ctx, cancel := context.WithTimeout(context.Background(), tt.inParams.ctxTimeout)
		defer cancel()
		conn, _ := NewConnection(ctx, tt.mongoHost)
		db := NewDatabase(tt.databaseName, conn)
		collection := NewCollection(tt.collectionName, db)
		docID, err := primitive.ObjectIDFromHex(tt.inParams.documentID)
		assert.Nil(t, err)
		got := collection.FindOne(ctx, bson.M{"_id": bson.M{"$eq": docID}})
		if tt.expectedErr {
			assert.Error(t, got.Err())
		} else {
			assert.Nil(t, got.Err())
		}
	}
}
