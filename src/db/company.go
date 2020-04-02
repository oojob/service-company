package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

	return "", err
}

// ReadCompany create company entity
func (db *Database) ReadCompany(id *primitive.ObjectID) (*model.Company, error) {
	companyCollection := db.Database("stayology").Collection("company")

	var company model.Company
	result := companyCollection.FindOne(context.Background(), &bson.M{"_id": id})
	if err := result.Decode(&company); err != nil {
		return nil, err
	}

	return &company, nil
}

// ReadCompanies create company entity
func (db *Database) ReadCompanies() (*mongo.Cursor, error) {
	companyCollection := db.Database("stayology").Collection("company")

	cursor, err := companyCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

// UpdateCompany create company entity
func (db *Database) UpdateCompany(id *primitive.ObjectID, in *bson.M) (string, error) {
	companyCollection := db.Database("stayology").Collection("company")

	result := companyCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, bson.M{"$set": in}, options.FindOneAndUpdate().SetReturnDocument(1))

	if result.Err() != nil {
		return "", result.Err()
	}

	var doc model.Company
	if err := result.Decode(&doc); err != nil {
		return "", err
	}

	return doc.ID.Hex(), nil
}

// DeleteCompany create company entity
func (db *Database) DeleteCompany(id *primitive.ObjectID) (string, error) {
	companyCollection := db.Database("stayology").Collection("company")

	var company model.Company
	result := companyCollection.FindOneAndDelete(context.Background(), bson.M{"_id": id})
	if err := result.Decode(&company); err != nil {
		return "", err
	}

	return company.ID.Hex(), nil
}
