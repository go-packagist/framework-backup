package foundation

import (
	"github.com/go-packagist/container"
	"reflect"
)

type ConcreteFunc func(*Application) interface{}

type binding struct {
	abstract string
	concrete ConcreteFunc
	shared   bool
}

// Application is the main application object.
type Application struct {
	basePath  string
	providers []Provider

	bindings  map[string]binding
	instances map[string]interface{}

	// services map[string]ServiceFunc

	*container.Container
}

// VERSION is the current version of the application.
const VERSION = "0.0.1"

// NewApplication creates a new application instance
func NewApplication(basePath string) *Application {
	app := &Application{
		basePath:  basePath,
		providers: []Provider{},
		bindings:  make(map[string]binding),
		instances: make(map[string]interface{}),
		// services:  make(map[string]ServiceFunc),
	}

	return app
}

// Register registers a provider with the application.
func (app *Application) Register(provider Provider) {
	if app.providerIsRegistered(provider) {
		return
	}

	provider.Register()

	app.providerMarkAsRegistered(provider)

	// todo
	// 1. bind
	// 2. boot
}

// providerIsRegistered return provider is registered
func (app *Application) providerIsRegistered(provider Provider) bool {
	for _, providerRegistered := range app.providers {
		if reflect.DeepEqual(providerRegistered, provider) {
			return true
		}
	}

	return false
}

// providerMarkAsRegistered provider mark as registered.
func (app *Application) providerMarkAsRegistered(provider Provider) {
	app.providers = append(app.providers, provider)
}

// GetProviders returns all registered providers.
func (app *Application) GetProviders() []Provider {
	return app.providers
}

// ------以下未完结版------

func (app *Application) Singleton(abstract string, concrete ConcreteFunc) {
	app.Bind(abstract, concrete, true)
}

func (app *Application) Bind(abstract string, concrete ConcreteFunc, shared bool) {
	app.bindings[abstract] = binding{
		abstract: abstract,
		concrete: concrete,
		shared:   shared,
	}
}

func (app *Application) Make(abstract string) interface{} {
	return app.Resolve(abstract)
}

func (app *Application) Resolve(abstract string) interface{} {
	// instance
	instance, ok := app.instances[abstract]
	if ok {
		return instance
	}

	// binding
	binding, ok2 := app.bindings[abstract]
	if !ok2 {
		panic(abstract + "not found")
	}

	// concrete(app)
	concrete := binding.concrete(app)

	if app.isShared(abstract) {
		app.instances[abstract] = concrete
	}

	return concrete
}

func (app *Application) isShared(abstract string) bool {
	return app.bindings[abstract].shared
}

// Singleton registers a service as a singleton.
// func (app *Application) Singleton(name string, service ServiceFunc) {
// 	app.services[name] = service
// }
//
// // GetService returns a registered service.
// func (app *Application) GetService(name string) interface{} {
// 	service, ok := app.services[name]
//
// 	if ok {
// 		return service(app)
// 	}
//
// 	panic("Service not found")
// }

// Version returns the current version of the application.
func (app *Application) Version() string {
	return VERSION
}
