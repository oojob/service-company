package db

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database base type struct
type Database struct {
	*mongo.Client
}

func New(config *Config) (*Database, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.DatabaseURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	return &Database{client}, nil
}
