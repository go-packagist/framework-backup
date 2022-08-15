package database

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/container"
)

type databaseProvider struct {
	container *container.Container
}

var _ container.Provider = (*databaseProvider)(nil)

func NewDatabaseProvider(c *container.Container) container.Provider {
	return &databaseProvider{
		container: c,
	}
}

func (p *databaseProvider) Register() {
	p.container.Singleton("database", func(c *container.Container) interface{} {
		return NewManager(c.MustMake("config").(*config.Config).Get("database").(*Config))
	})
}
