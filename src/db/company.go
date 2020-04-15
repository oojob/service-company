package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/oojob/service-company/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateCompany create company entity
func (db *Database) CreateCompany(in *model.Company) (string, error) {
	var inerstionID string
	oojob := db.Database("stayology")
	client := oojob.Client()
	companyCollection := oojob.Collection("company")

	// start the session
	session, err := client.StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(context.Background())

	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result, err := companyCollection.InsertOne(sessionContext, in)
		if err != nil {
			return "", err
		}

		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			inerstionID = oid.Hex()
		}

		return "", nil
	})

	return inerstionID, err
}

// ReadCompany create company entity
// should be modified by dibya
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

// ReadAllCompanies read all company entities
func (db *Database) ReadAllCompanies(skip string, limit int64) (*[]*model.Company, error) {
	companyCollection := db.Database("stayology").Collection("company")

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	// store all companies result
	var companies []*model.Company

	cursor, err := companyCollection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		// create an empty value to hold single decoded company value
		var company model.Company
		if err := cursor.Decode(&company); err != nil {
			return nil, err
		}
		// To get the raw bson bytes use cursor.Current
		// raw := cursor.Current
		companies = append(companies, &company)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())
	return &companies, nil
}
