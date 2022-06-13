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
	hashedValue := app.MustMake("hash").(*Manager).Driver().MustMake(value)
	assert.True(t, app.MustMake("hash").(*Manager).Driver().Check(value, hashedValue))

	// hasher.bcrypt
	hashedValue2 := app.MustMake("hasher.bcrypt").(*BcryptHasher).MustMake(value)
	assert.True(t, app.MustMake("hasher.bcrypt").(*BcryptHasher).Check(value, hashedValue2))

	// hasher.md5
	hashedValue3 := app.MustMake("hasher.md5").(*Md5Hasher).MustMake(value)
	assert.True(t, app.MustMake("hasher.md5").(*Md5Hasher).Check(value, hashedValue3))
}

func createApp() *foundation.Application {
	app := foundation.NewApplication("./")

	app.Register(config.NewConfigProvider(app))

	// set config
	app.MustMake("config").(*config.Config).Set("hashing", &Config{
		Driver: "bcrypt",
	})

	return app
}
