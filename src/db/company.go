package db

import (
	"context"
	"errors"

	"github.com/oojob/company/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateCompany create company entity
func (db *Database) CreateCompany(in *model.Company) (string, error) {
	companyCollection := db.Database("stayology").Collection("company")

	result, err := companyCollection.InsertOne(context.Background(), in)

	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", errors.New("invalid id")
}
