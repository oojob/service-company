package api

import (
	"context"

	"github.com/oojob/company/src/model"
	company "github.com/oojob/protorepo-company-go"
)

// CreateCompany function for creating company
func (c *API) CreateCompany(ctx context.Context, in *company.CreateCompanyReq) (*company.CreateCompanyRes, error) {

	companyData := model.Company{
		Name:        in.GetName(),
		Description: "dodo duck lives here",
		CreatedBy:   "1234567",
		URL:         "http://dododuck.io",
		Logo:        "https://azure.upload.net/dododuck.png",
		Location:    "Moscos",
		FoundedYear: 2020,
		NoOfEmployees: model.NoOfEmployees{
			2,
			4,
		},
	}

	context := c.App.NewContext()
	result, err := context.CreateCompany(&companyData)

	if err != nil {
		return nil, err
	}

	return &company.CreateCompanyRes{Status: true, Id: result}, nil
}
