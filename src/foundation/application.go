package foundation

import "github.com/go-packagist/container"

type Application struct {
	basePath string

	services map[string]interface{}

	*container.Container

	providers map[string]Provider
}

const VERSION = "0.0.1"

func NewApplication(basePath string) *Application {
	app := &Application{
		basePath:  basePath,
		providers: make(map[string]Provider),
		services:  make(map[string]interface{}),
	}

	return app
}

func (app *Application) Register(name string, provider Provider) Provider {
	app.providers[name] = provider

	provider.Register()

	return provider
}

func (app *Application) GetProvider(name string) Provider {
	return app.providers[name]
}

func (app *Application) GetProviders() map[string]Provider {
	return app.providers
}

func (app *Application) GetService(name string) interface{} {
	service, ok := app.services[name]

	if ok {
		return service
	}

	panic("Service not found")

	return nil
}

func (app *Application) Version() string {
	return VERSION
}
