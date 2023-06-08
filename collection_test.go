package mongoose

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_NewCollection(t *testing.T) {
	conn, _ := NewConnection(context.Background(), "mongodb://mock-string")
	db := NewDatabase("mock-db", conn)
	got := NewCollection("mock-collection", db)
	assert.NotEmpty(t, got)
}

func Test_CollectionHelper_InsertOne(t *testing.T) {
	type inParams struct {
		ctx      context.Context
		document interface{}
	}
	tests := []struct {
		name        string
		expectedErr bool
		inParams    inParams
	}{
		{
			name:        "Test insert one with error retuning",
			expectedErr: true,
			inParams: inParams{
				ctx:      context.Background(),
				document: bson.M{"mock": "document"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, _ := NewConnection(context.Background(), "mongodb://mock-string")
			db := NewDatabase("mock-db", conn)
			collection := NewCollection("mock-collection", db)
			got, err := collection.InsertOne(tt.inParams.ctx, tt.inParams.document)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NotNil(t, got)
			}
		})
	}

}
