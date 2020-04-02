package app

import (
	"github.com/oojob/company/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateCompany creates a company
func (ctx *Context) CreateCompany(company *model.Company) (string, error) {
	return ctx.Database.CreateCompany(company)
}

// ReadCompany creates a company
func (ctx *Context) ReadCompany(id *primitive.ObjectID) (*model.Company, error) {
	return ctx.Database.ReadCompany(id)
}

// ReadCompanies creates a company
func (ctx *Context) ReadCompanies() (*mongo.Cursor, error) {
	return ctx.Database.ReadCompanies()
}

// UpdateCompany creates a company
func (ctx *Context) UpdateCompany(id *primitive.ObjectID, company *bson.M) (string, error) {
	return ctx.Database.UpdateCompany(id, company)
}

// DeleteCompany creates a company
func (ctx *Context) DeleteCompany(id *primitive.ObjectID) (string, error) {
	return ctx.Database.DeleteCompany(id)
}
