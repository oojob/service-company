package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NoOfEmployees create no. of employees range
type NoOfEmployees struct {
	Min int64 `bson:"min,omitempty"`
	Max int64 `bson:"max,omitempty"`
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
	FoundedYear   string             `bson:"founded_year,omitempty"`
	LastActive    time.Time          `bson:"last_active,omitempty"`
	HiringStatus  bool               `bosn:"hiring_status,omitempty"`
	Skills        []string           `bson:"skills,omitempty"`
	NoOfEmployees NoOfEmployees      `bson:"no_of_employees,omitempty"`

	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}
