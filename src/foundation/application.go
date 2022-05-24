package foundation

import "github.com/go-packagist/container"

// ServiceFunc is a function that returns a service.
type ServiceFunc func(app *Application) interface{}

// Application is the main application object.
type Application struct {
	basePath string

	services map[string]ServiceFunc

	*container.Container

	providers map[string]Provider
}

// VERSION is the current version of the application.
const VERSION = "0.0.1"

// NewApplication creates a new application instance
func NewApplication(basePath string) *Application {
	app := &Application{
		basePath:  basePath,
		providers: make(map[string]Provider),
		services:  make(map[string]ServiceFunc),
	}

	return app
}

// Register registers a provider with the application.
func (app *Application) Register(name string, provider Provider) Provider {
	app.providers[name] = provider

	provider.Register()

	return provider
}

// GetProvider returns a registered provider.
func (app *Application) GetProvider(name string) Provider {
	return app.providers[name]
}

// GetProviders returns all registered providers.
func (app *Application) GetProviders() map[string]Provider {
	return app.providers
}

// Singleton registers a service as a singleton.
func (app *Application) Singleton(name string, service ServiceFunc) {
	app.services[name] = service
}

// GetService returns a registered service.
func (app *Application) GetService(name string) interface{} {
	service, ok := app.services[name]

	if ok {
		return service(app)
	}

	panic("Service not found")
}

// Version returns the current version of the application.
func (app *Application) Version() string {
	return VERSION
}
