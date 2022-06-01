package hashing

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
)

// hashProvider represents the hashing provider.
type hashProvider struct {
	app *foundation.Application
}

var _ foundation.Provider = (*hashProvider)(nil)

// NewHashProvider Bootstrap bootstraps the hashing services.
func NewHashProvider(app *foundation.Application) *hashProvider {
	return &hashProvider{
		app: app,
	}
}

// Register registers the hashing services into the application.
func (p *hashProvider) Register() {
	p.app.Singleton("hash", func(app *foundation.Application) interface{} {
		return NewHashManager(app.MustMake("config").(*config.Config).Get("hashing").(map[string]interface{}))
	})

	p.app.Singleton("hasher.bcrypt", func(app *foundation.Application) interface{} {
		return NewBcryptHasher()
	})
}
