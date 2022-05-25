package foundation

import (
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
}

// VERSION is the current version of the application.
const VERSION = "0.0.1"

var instance *Application

// NewApplication creates a new application instance
func NewApplication(basePath string) *Application {
	app := &Application{
		basePath:  basePath,
		providers: []Provider{},
		bindings:  make(map[string]binding),
		instances: make(map[string]interface{}),
		// services:  make(map[string]ServiceFunc),
	}

	// 设置常驻变量
	SetInstance(app)

	return app
}

func SetInstance(app *Application) {
	instance = app
}

func GetInstance() *Application {
	return instance
}

func Instance() *Application {
	return GetInstance()
}

func App() *Application {
	return GetInstance()
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

// Singleton Register a shared binding in the container.
func (app *Application) Singleton(abstract string, concrete ConcreteFunc) {
	app.Bind(abstract, concrete, true)
}

// Bind Register a binding with the container.
func (app *Application) Bind(abstract string, concrete ConcreteFunc, shared bool) {
	app.bindings[abstract] = binding{
		abstract: abstract,
		concrete: concrete,
		shared:   shared,
	}
}

// Make Resolve the given type from the container.
func (app *Application) Make(abstract string) interface{} {
	return app.Resolve(abstract)
}

// Resolve the given type from the container.
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

// isShared Determine if a given type is shared.
func (app *Application) isShared(abstract string) bool {
	return app.bindings[abstract].shared
}

// Version returns the current version of the application.
func (app *Application) Version() string {
	return VERSION
}
