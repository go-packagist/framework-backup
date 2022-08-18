package config

import (
	"github.com/go-packagist/framework/container"
)

type configProvider struct {
	container *container.Container
}

var _ container.Provider = (*configProvider)(nil)

func NewConfigProvider(c *container.Container) container.Provider {
	return &configProvider{
		container: c,
	}
}

func (p *configProvider) Register() {
	options := &Options{}

	if basePath, err := p.container.Get("path.base"); err == nil {
		options.EnvPath = basePath.(string)
	}

	p.container.Instance("config", New(options))
}

func (p *configProvider) Boot() {

}
