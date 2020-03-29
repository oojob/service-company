package api

import (
	"github.com/oojob/company/src/app"
)

// CompanyServer base server struct
type CompanyServer struct{}

// CompanyAPI api base struct
type CompanyAPI struct {
	App           *app.App
	Config        *Config
	CompanyServer CompanyServer
}

// New new api instance
func New(a *app.App) (api *CompanyAPI, err error) {
	api = &CompanyAPI{App: a}

	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	companyServer := CompanyServer{}
	api.CompanyServer = companyServer

	return api, nil
}
