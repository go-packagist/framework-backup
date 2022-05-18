package foundation

import "github.com/go-packagist/container"

type Application struct {
	basePath string

	*container.Container

	providers map[string]Provider
}

const VERSION = "0.0.1"

func NewApplication(basePath string) *Application {
	app := &Application{
		basePath:  basePath,
		providers: make(map[string]Provider),
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

func (app *Application) Version() string {
	return VERSION
}
