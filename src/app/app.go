package app

// App Application
type App struct {
	Config *Config
}

// New Application new instance
func New() (app *App, err error) {
	app = &App{}

	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	return app, err
}
