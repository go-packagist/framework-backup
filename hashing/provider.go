package hashing

import (
	"github.com/go-packagist/framework/container"
	"github.com/go-packagist/framework/contracts/provider"
)

// HashProvider represents the hashing provider.
type HashProvider struct {
	container *container.Container
}

var _ provider.Provider = (*HashProvider)(nil)

// NewHashProvider Bootstrap bootstraps the hashing services.
func NewHashProvider(c *container.Container) *HashProvider {
	return &HashProvider{
		container: c,
	}
}

// Register registers the hashing services into the application.
func (p *HashProvider) Register() {
	p.container.Singleton("hash", func(c *container.Container) interface{} {
		return NewManager(&Config{
			Driver: "md5", // todo: modify to config
		})
	})
}

func (p *HashProvider) Boot() {

}
