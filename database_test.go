package mongoose

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_NewDatabase(t *testing.T) {
	tests := []struct {
		name     string
		inDbName string
	}{
		{
			name:     "Test new DB with valid name 'mock-name1'",
			inDbName: "mock-name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, _ := NewConnection(context.Background(), "mongodb://mock-string")
			got := NewDatabase(tt.inDbName, conn)
			assert.NotEmpty(t, got)
		})
	}
}

func Test_DatabaseHelper_Client(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Get valid client",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, _ := NewConnection(context.Background(), "mongodb://mock-string")
			db := NewDatabase("mock-databse", conn)
			got := db.Client()
			assert.NotEmpty(t, got)

		})
	}
}

func Test_DatabaseHelper_Collection(t *testing.T) {
	tests := []struct {
		name             string
		inCollectionName string
	}{
		{
			name:             "Get valid client",
			inCollectionName: "mock-collection",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, _ := NewConnection(context.Background(), "mongodb://mock-string")
			db := NewDatabase("mock-databse", conn)
			got := db.Collection(tt.inCollectionName)
			assert.NotEmpty(t, got)

		})
	}
}

func Test_DatabaseHelper_UseSession(t *testing.T) {
	tests := []struct {
		name          string
		inFunc        func(mongo.SessionContext) error
		expecterError error
	}{
		{
			name: "Test function no error",
			inFunc: func(sc mongo.SessionContext) error {
				return nil
			},
		},
		{
			name: "Test function that returns an error",
			inFunc: func(sc mongo.SessionContext) error {
				return errors.New("mock-error")
			},
			expecterError: errors.New("mock-error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, _ := NewConnection(context.Background(), "mongodb://mock-string")
			db := NewDatabase("mock-databse", conn)
			got := db.UseSession(context.Background(), tt.inFunc)
			if tt.expecterError != nil {
				assert.Error(t, got)
				assert.Equal(t, tt.expecterError.Error(), got.Error())
			} else {
				assert.Nil(t, got)
			}

		})
	}
}
