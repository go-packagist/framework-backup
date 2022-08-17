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

	app.Register(config.NewConfigProvider(app.Container))

	// set config-hasing
	app.MustMake("config").(*config.Config).Set("hashing", &hashing.Config{
		Driver: "bcrypt",
	})

	return app
}

func TestFacade_App(t *testing.T) {
	createApp()

	// assert.Equal(t, "./app", App().getPath())
	assert.Equal(t, "./app", App().Container.MustMake("path").(string))
}

func TestFacades_Hashing(t *testing.T) {
	app := createApp()
	app.Register(hashing.NewHashProvider(app.Container))

	value := "123456"

	hashedValue1 := MustHash().Driver().MustMake(value)
	assert.True(t, MustHash().Driver().Check(value, hashedValue1))
	hashedValue2 := MustBcryptHasher().MustMake(value)
	assert.True(t, MustBcryptHasher().Check(value, hashedValue2))
}

func TestFacades_Config(t *testing.T) {
	createApp()

	assert.Equal(t, "bcrypt", MustConfig().Get("hashing").(*hashing.Config).Driver)
}
