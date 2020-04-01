package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type NoOfEmployees struct {
	Min int `bson:"min,omitempty"`
	Max int `bson:"max,omitempty"`
}

// Company base entity
type Company struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name,omitempty"`
	Description   string             `bson:"description,omitempty"`
	CreatedBy     string             `bson:"created_by,omitempty"`
	URL           string             `bson:"url,omitempty"`
	Logo          string             `bson:"logo,omitempty"`
	Location      string             `bson:"location,omitempty"`
	FoundedYear   int                `bson:"founded_year,omitempty"`
	NoOfEmployees NoOfEmployees
}
