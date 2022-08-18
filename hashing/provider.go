package hashing

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/container"
)

// hashProvider represents the hashing provider.
type hashProvider struct {
	container *container.Container
}

var _ container.Provider = (*hashProvider)(nil)

// NewHashProvider Bootstrap bootstraps the hashing services.
func NewHashProvider(c *container.Container) *hashProvider {
	return &hashProvider{
		container: c,
	}
}

// Register registers the hashing services into the application.
func (p *hashProvider) Register() {
	p.container.Singleton("hash", func(c *container.Container) interface{} {
		return NewManager(c.MustMake("config").(*config.Config).Get("hashing").(*Config))
	})

	p.container.Singleton("hasher.bcrypt", func(c *container.Container) interface{} {
		return NewBcryptHasher()
	})

	p.container.Singleton("hasher.md5", func(c *container.Container) interface{} {
		return NewMd5Hasher()
	})
}

func (p *hashProvider) Boot() {

}
