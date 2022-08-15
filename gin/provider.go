package gin

import (
	"github.com/go-packagist/framework/container"
)

type Provider struct {
	container *container.Container
}

var _ container.Provider = (*Provider)(nil)

func NewGinProvider(c *container.Container) container.Provider {
	return &Provider{
		container: c,
	}
}

func (p Provider) Register() {
	p.container.Singleton("gin", func(c *container.Container) interface{} {
		return NewGin(c)
	})
}
