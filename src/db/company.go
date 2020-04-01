package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

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

// ReadCompany create company entity
func (db *Database) ReadCompany(in string) (*model.Company, error) {
	companyCollection := db.Database("stayology").Collection("company")

	id, err := primitive.ObjectIDFromHex(in)
	if err != nil {
		return nil, err
	}

	var company model.Company
	err = companyCollection.FindOne(context.Background(), model.Company{ID: id}).Decode(&company)

	return &company, nil
}

// ReadCompanies create company entity
func (db *Database) ReadCompanies() (*mongo.Cursor, error) {
	companyCollection := db.Database("stayology").Collection("company")

	cursor, err := companyCollection.Find(context.Background(), model.Company{})
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

// UpdateCompany create company entity
func (db *Database) UpdateCompany(in *model.Company) (string, error) {
	companyCollection := db.Database("stayology").Collection("company")

	result := companyCollection.FindOneAndUpdate(context.Background(), model.Company{ID: in.ID}, in, nil)

	if result.Err() != nil {
		return "", result.Err()
	}

	var doc model.Company
	decodeErr := result.Decode(&doc)
	if decodeErr != nil {
		return "", decodeErr
	}

	return doc.ID.Hex(), nil
}

// DeleteCompany create company entity
func (db *Database) DeleteCompany(in string) (string, error) {
	companyCollection := db.Database("stayology").Collection("company")

	id, err := primitive.ObjectIDFromHex(in)
	if err != nil {
		return "", err
	}

	var company model.Company
	err = companyCollection.FindOneAndDelete(context.Background(), model.Company{ID: id}).Decode(&company)
	if err != nil {
		return "", err
	}

	return company.ID.Hex(), nil
}
