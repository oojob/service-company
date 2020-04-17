package api

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang/protobuf/ptypes"

	company "github.com/oojob/protorepo-company-go"
	"github.com/oojob/service-company/src/model"
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
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid Document: %v", err),
		)
	}

	return &company.Id{Id: result}, nil
}

// ReadCompany read company data
func (c *API) ReadCompany(ctx context.Context, in *company.Id) (*company.Company, error) {
	context := c.App.NewContext()

	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	result, err := context.ReadCompany(&id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Object Id %s: %v", in.GetId(), err))
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

	return companyResponse, nil
}

// ReadCompanies read all companies data
func (c *API) ReadCompanies(in *company.Empty, stream company.CompanyService_ReadCompaniesServer) error {
	ctx := c.App.NewContext()
	result := &model.Company{}

	cursor, err := ctx.ReadCompanies()
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
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
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err))
	}

	companyData := bson.M{
		"name":         in.GetName(),
		"description":  in.GetDescription(),
		"created_by":   in.GetCreatedBy(),
		"url":          in.GetUrl(),
		"logo":         in.GetLogo(),
		"location":     in.GetLocation(),
		"founded_year": in.GetFoundedYear(),
		"hiringstatus": in.GetHiringStatus(),
		"no_of_employees": bson.M{
			"min": in.NoOfEmployees.Min,
			"max": in.NoOfEmployees.Max,
		},
		"skills":      in.GetSkills(),
		"last_active": time.Now(),
		"updated_at":  time.Now(),
	}

	context := c.App.NewContext()
	result, err := context.UpdateCompany(&id, &companyData)

	if err != nil {
		return nil, err
	}

	return &company.Id{Id: result}, nil
}

// DeleteCompany delete company adat
func (c *API) DeleteCompany(ctx context.Context, in *company.Id) (*company.Id, error) {
	context := c.App.NewContext()

	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	result, err := context.DeleteCompany(&id)

	if err != nil {
		return nil, err
	}

	return &company.Id{Id: result}, nil
}

// ReadAllCompanies data
func (c *API) ReadAllCompanies(ctx context.Context, in *company.Pagination) (*company.CompanyAllResponse, error) {
	context := c.App.NewContext()
	limit := in.GetLimit()
	skip := in.GetSkip()

	response, error := context.ReadAllCompanies(skip, limit)
	if error != nil {
		return nil, error
	}

	var companies []*company.Company
	for _, result := range *response {
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

		company := &company.Company{
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

		companies = append(companies, company)
	}

	return &company.CompanyAllResponse{
		Companies: companies,
	}, nil
}

// Check check the context
func (c *API) Check(ctx context.Context, in *company.HealthCheckRequest) (*company.HealthCheckResponse, error) {
	return &company.HealthCheckResponse{
		Status: company.HealthCheckResponse_SERVING,
	}, nil
}

// Watch watch the serving status
func (c *API) Watch(_ *company.HealthCheckRequest, stream company.CompanyService_WatchServer) error {
	stream.Send(&company.HealthCheckResponse{
		Status: company.HealthCheckResponse_SERVING,
	})

	return nil
}
