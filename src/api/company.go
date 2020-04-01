package api

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/oojob/company/src/model"
	company "github.com/oojob/protorepo-company-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCompany function for creating company
func (c *API) CreateCompany(ctx context.Context, in *company.Company) (*company.Id, error) {

	companyData := model.Company{
		Name:         in.GetName(),
		Description:  in.GetDescription(),
		CreatedBy:    in.GetCreatedBy(),
		URL:          in.GetUrl(),
		Logo:         in.GetLogo(),
		Location:     in.GetLocation(),
		FoundedYear:  in.GetFoundedYear(),
		HiringStatus: in.GetHiringStatus(),
		NoOfEmployees: model.NoOfEmployees{
			Min: in.GetNoOfEmployees().Min,
			Max: in.GetNoOfEmployees().Max,
		},
		Skills:     in.GetSkills(),
		LastActive: time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	context := c.App.NewContext()
	result, err := context.CreateCompany(&companyData)

	if err != nil {
		return nil, err
	}

	return &company.Id{Id: result}, nil
}

// ReadCompany read company data
func (c *API) ReadCompany(ctx context.Context, in *company.Id) (*company.Company, error) {
	context := c.App.NewContext()
	result, err := context.ReadCompany(in.GetId())

	if err != nil {
		return nil, err
	}

	createdAt, err := ptypes.TimestampProto(result.CreatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}

	updatedAt, err := ptypes.TimestampProto(result.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}

	lastActive, err := ptypes.TimestampProto(result.LastActive)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}

	return &company.Company{
		Id:           result.ID.Hex(),
		Name:         result.Name,
		Description:  result.Description,
		CreatedBy:    result.CreatedBy,
		Url:          result.URL,
		Logo:         result.Logo,
		Location:     result.Location,
		FoundedYear:  result.FoundedYear,
		HiringStatus: result.HiringStatus,
		Skills:       result.Skills,
		NoOfEmployees: &company.Range{
			Min: result.NoOfEmployees.Min,
			Max: result.NoOfEmployees.Max,
		},
		LastActive: lastActive,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}, nil
}

// ReadCompanies read all companies data
func (c *API) ReadCompanies(in *company.Empty, stream company.CompanyService_ReadCompaniesServer) error {
	ctx := c.App.NewContext()
	result := &model.Company{}

	cursor, err := ctx.ReadCompanies()
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		err := cursor.Decode(result)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		createdAt, err := ptypes.TimestampProto(result.CreatedAt)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		updatedAt, err := ptypes.TimestampProto(result.UpdatedAt)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		lastActive, err := ptypes.TimestampProto(result.LastActive)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		companyResponse := &company.Company{
			Id:           result.ID.Hex(),
			Name:         result.Name,
			Description:  result.Description,
			CreatedBy:    result.CreatedBy,
			Url:          result.URL,
			Logo:         result.Logo,
			Location:     result.Location,
			FoundedYear:  result.FoundedYear,
			HiringStatus: result.HiringStatus,
			Skills:       result.Skills,
			NoOfEmployees: &company.Range{
				Min: result.NoOfEmployees.Min,
				Max: result.NoOfEmployees.Max,
			},
			LastActive: lastActive,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		}

		stream.Send(companyResponse)
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}

	return nil
}

// UpdateCompany update company data
func (c *API) UpdateCompany(ctx context.Context, in *company.Company) (*company.Id, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, err
	}

	companyData := model.Company{
		ID:           id,
		Name:         in.GetName(),
		Description:  in.GetDescription(),
		CreatedBy:    in.GetCreatedBy(),
		URL:          in.GetUrl(),
		Logo:         in.GetLogo(),
		Location:     in.GetLocation(),
		FoundedYear:  in.GetFoundedYear(),
		HiringStatus: in.GetHiringStatus(),
		NoOfEmployees: model.NoOfEmployees{
			Min: in.GetNoOfEmployees().Min,
			Max: in.GetNoOfEmployees().Max,
		},
		Skills:     in.GetSkills(),
		LastActive: time.Now(),
		UpdatedAt:  time.Now(),
	}

	context := c.App.NewContext()
	result, err := context.UpdateCompany(&companyData)

	if err != nil {
		return nil, err
	}

	return &company.Id{Id: result}, nil
}

// DeleteCompany delete company adat
func (c *API) DeleteCompany(ctx context.Context, in *company.Id) (*company.Id, error) {
	context := c.App.NewContext()
	result, err := context.DeleteCompany(in.GetId())

	if err != nil {
		return nil, err
	}

	return &company.Id{Id: result}, nil
}
