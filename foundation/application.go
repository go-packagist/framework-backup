package foundation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
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

	rwlock *sync.RWMutex
}

// VERSION is the current version of the application.
const VERSION = "0.0.1"

var instance *Application

// NewApplication creates a new application instance
func NewApplication(basePath string) *Application {
	app := &Application{
		providers: []Provider{},
		bindings:  make(map[string]binding),
		instances: make(map[string]interface{}),
		rwlock:    &sync.RWMutex{},
	}

	// 设置basePath
	app.setBasePath(basePath)

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
	app.rwlock.Lock()
	defer app.rwlock.Unlock()

	app.bindings[abstract] = binding{
		abstract: abstract,
		concrete: concrete,
		shared:   shared,
	}
}

func (app *Application) Instance(abstract string, concrete interface{}) {
	app.rwlock.Lock()
	defer app.rwlock.Unlock()

	app.instances[abstract] = concrete
}

// Make Resolve the given type from the container.
func (app *Application) Make(abstract string) (interface{}, error) {
	return app.Resolve(abstract)
}

// MustMake Resolve the given type from the container or panic.
func (app *Application) MustMake(abstract string) interface{} {
	concrete, err := app.Make(abstract)
	if err != nil {
		panic(err)
	}

	return concrete
}

// Resolve the given type from the container.
func (app *Application) Resolve(abstract string) (interface{}, error) {
	// instance
	instance, ok := app.instances[abstract]
	if ok {
		return instance, nil
	}

	// binding
	binding, ok2 := app.bindings[abstract]
	if !ok2 {
		return nil, errors.New(fmt.Sprintf("[%s] binding not found", abstract))
	}

	// concrete(app)
	concrete := binding.concrete(app)

	if app.isShared(abstract) {
		app.Instance(abstract, concrete)
	}

	return concrete, nil
}

// isShared Determine if a given type is shared.
func (app *Application) isShared(abstract string) bool {
	return app.bindings[abstract].shared
}

// SetBasePath sets the base path for the application.
func (app *Application) setBasePath(basePath string) {
	app.basePath = strings.TrimRight(basePath, "/")

	app.bindPathInApplication()
}

// bindPathInApplication bind path in application
func (app *Application) bindPathInApplication() {
	app.Instance("path", app.getPath())
	app.Instance("path.base", app.getBasePath())
	app.Instance("path.config", app.getConfigPath())
	app.Instance("path.bootstrap", app.getBootstrapPath())
}

// getPath returns the app path to the base of the application.
func (app *Application) getPath() string {
	return app.basePath + "/" + "app"
}

// getBasePath returns the base path of the application.
func (app *Application) getBasePath() string {
	return app.basePath
}

// getConfigPath returns the path to the config folder.
func (app *Application) getConfigPath() string {
	return app.basePath + "/" + "config"
}

// getBootstrapPath returns the path to the bootstrap folder.
func (app *Application) getBootstrapPath() string {
	return app.basePath + "/" + "bootstrap"
}

// Version returns the current version of the application.
func (app *Application) Version() string {
	return VERSION
}
