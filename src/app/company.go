package app

import "github.com/oojob/company/src/model"

// CreateCompany creates a company
func (ctx *Context) CreateCompany(company *model.Company) (string, error) {
	return ctx.Database.CreateCompany(company)
}
