package database

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
)

type databaseProvider struct {
	app *foundation.Application
}

var _ foundation.Provider = (*databaseProvider)(nil)

func NewDatabaseProvider(app *foundation.Application) foundation.Provider {
	return &databaseProvider{
		app: app,
	}
}

func (p *databaseProvider) Register() {
	p.app.Singleton("database", func(app *foundation.Application) interface{} {
		return NewManager(app.MustMake("config").(*config.Config).Get("database").(*Config))
	})
}
