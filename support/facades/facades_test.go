package facades

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/hashing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createApp() *foundation.Application {
	app := foundation.NewApplication("./")

	app.Register(config.NewConfigProvider(app))

	// set config-hasing
	app.MustMake("config").(*config.Config).Set("hashing", map[string]interface{}{
		"driver": "bcrypt",
	})

	return app
}

func TestFacades_App(t *testing.T) {
	app := createApp()

	assert.Equal(t, app, App())
}

func TestFacades_Config(t *testing.T) {
	createApp()

	assert.Equal(t, "bcrypt", MustConfig().Get("hashing.driver"))
}

func TestFacades_Hashing(t *testing.T) {
	app := createApp()
	app.Register(hashing.NewHashProvider(app))

	value := "123456"

	hashedValue1 := MustHash().Driver().MustMake(value)
	assert.True(t, MustHash().Driver().Check(value, hashedValue1))
	hashedValue2 := MustBcryptHasher().MustMake(value)
	assert.True(t, MustBcryptHasher().Check(value, hashedValue2))
}
