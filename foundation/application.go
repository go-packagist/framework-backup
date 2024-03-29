package foundation

import (
	"github.com/go-packagist/framework/container"
	"github.com/go-packagist/framework/version"
)

// Application is the main application object.
type Application struct {
	*container.Container
	booted bool
}

// Application is the main application object.
var instance *Application

// NewApplication creates a new application instance
func NewApplication() *Application {
	app := &Application{
		Container: container.New(),
		booted:    false,
	}

	// register the application instance
	app.Container.Instance("app", app)

	// set the application instance
	SetInstance(app)

	return app
}

// SetInstance sets the application instance.
func SetInstance(app *Application) {
	instance = app
}

// GetInstance returns the application instance.
func GetInstance() *Application {
	return instance
}

// Instance returns the application instance.
func Instance() *Application {
	return GetInstance()
}

// App returns the application instance.
func App() *Application {
	return GetInstance()
}

// Boot the application.
func (app *Application) Boot() {
	if app.booted {
		return
	}

	for _, provider := range app.Container.GetProviders() {
		provider.Boot()
	}

	app.booted = true
}

// Version returns the current version of the application.
func (app *Application) Version() string {
	return version.Version
}
