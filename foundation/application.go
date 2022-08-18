package foundation

import (
	"github.com/go-packagist/framework/container"
	"strings"
)

// Application is the main application object.
type Application struct {
	basePath  string
	Container *container.Container
}

// VERSION is the current version of the application.
const VERSION = "0.0.1"

var instance *Application

// NewApplication creates a new application instance
func NewApplication(basePath string) *Application {
	app := &Application{
		Container: container.NewContainer(),
	}

	// 注入容器
	app.Container.Instance("app", app)

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

func (app *Application) Register(provider container.Provider) {
	app.Container.Register(provider)
}

func (app *Application) Instance(abstract string, concrete interface{}) {
	app.Container.Instance(abstract, concrete)
}

// Singleton Register a shared binding in the container.
func (app *Application) Singleton(abstract string, concrete container.ConcreteFunc) {
	app.Container.Singleton(abstract, concrete)
}

func (app *Application) Bind(abstract string, concrete container.ConcreteFunc, shared bool) {
	app.Container.Bind(abstract, concrete, shared)
}

// Make Resolve the given type from the container.
func (app *Application) Make(abstract string) (interface{}, error) {
	return app.Container.Make(abstract)
}

// MustMake Resolve the given type from the container or panic.
func (app *Application) MustMake(abstract string) interface{} {
	concrete, err := app.Make(abstract)

	if err != nil {
		panic(err)
	}

	return concrete
}

// SetBasePath sets the base path for the application.
func (app *Application) setBasePath(basePath string) {
	app.basePath = strings.TrimRight(basePath, "/")

	app.bindPathInApplication()
}

// bindPathInApplication bind path in application
func (app *Application) bindPathInApplication() {
	app.Container.Instance("path", app.getPath())
	app.Container.Instance("path.base", app.getBasePath())
	app.Container.Instance("path.config", app.getConfigPath())
	app.Container.Instance("path.bootstrap", app.getBootstrapPath())
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
