package api

import (
	"context"

	company "github.com/oojob/protorepo-company-go"
)

// CreateCompany function for creating company
func (s *CompanyServer) CreateCompany(ctx context.Context, in *company.CreateCompanyReq) (*company.CreateCompanyRes, error) {

	return &company.CreateCompanyRes{Status: true, Id: "12345"}, nil
}
