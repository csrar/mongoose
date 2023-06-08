package mongoose

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConnection(t *testing.T) {
	type args struct {
		ctx              context.Context
		connectionString string
	}
	tests := []struct {
		name        string
		inParams    args
		expectedErr string
	}{
		{
			name: "Test valid connection string",
			inParams: args{
				ctx:              context.Background(),
				connectionString: "mongodb://mock-string",
			},
		}, {
			name: "Test invalid connetion string",
			inParams: args{
				ctx:              context.Background(),
				connectionString: "mock-string",
			},
			expectedErr: "error parsing uri: scheme must be \"mongodb\" or \"mongodb+srv\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConnection(tt.inParams.ctx, tt.inParams.connectionString)
			if tt.expectedErr != "" {
				assert.Equal(t, tt.expectedErr, err.Error())
			} else {
				assert.NotEmpty(t, got)
			}
		})
	}
}

func Test_ClientHelper_Ping(t *testing.T) {
	tests := []struct {
		name        string
		hostname    string
		expectedErr bool
	}{
		{
			name:        "Test ping with invalid host",
			hostname:    "mongodb://mock-string",
			expectedErr: true,
		},
	}
	for _, tt := range tests {
		conn, _ := NewConnection(context.Background(), tt.hostname)
		got := conn.Ping(context.Background())
		if tt.expectedErr {
			assert.Error(t, got)
		} else {
			assert.Nil(t, got)
		}
	}
}

func Test_clientHelper_Database(t *testing.T) {
	conn, _ := NewConnection(context.Background(), "mongodb://mock-string")
	got := conn.Database("mock-database")
	assert.NotEmpty(t, got)
}
