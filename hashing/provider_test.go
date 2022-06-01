package hashing

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	app := createApp()

	app.Register(NewHashProvider(app))

	// hash
	value := "123456"
	hashedValue := app.MustMake("hash").(*HashManager).Driver().MustMake(value)
	assert.True(t, app.MustMake("hash").(*HashManager).Driver().Check(value, hashedValue))

	// hash.bcrypt
	hashedValue2 := app.MustMake("hash.bcrypt").(*BcryptHasher).MustMake(value)
	assert.True(t, app.MustMake("hash.bcrypt").(*BcryptHasher).Check(value, hashedValue2))

	// Facade
	hashedValue3 := FacadeMustHash().Driver().MustMake(value)
	assert.True(t, FacadeMustHash().Driver().Check(value, hashedValue3))
	hashedValue4 := FacadeMustHashBcrypt().MustMake(value)
	assert.True(t, FacadeMustHashBcrypt().Check(value, hashedValue4))
}

func createApp() *foundation.Application {
	app := foundation.NewApplication("./")

	app.Register(config.NewConfigProvider(app))

	// set config
	app.MustMake("config").(*config.Config).Set("hashing", map[string]interface{}{
		"driver": "bcrypt",
	})

	return app
}
