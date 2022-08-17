package config

import (
	"github.com/go-packagist/framework/container"
)

type Provider struct {
	container *container.Container
}

var _ container.Provider = (*Provider)(nil)

func NewConfigProvider(c *container.Container) container.Provider {
	return &Provider{
		container: c,
	}
}

func (p *Provider) Register() {
	p.container.Instance("config", New(&Options{}))
}
