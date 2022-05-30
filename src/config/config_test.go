package config

import (
	"fmt"
	"github.com/go-packagist/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_New(t *testing.T) {
	config := New(&Options{
		EnvPath: "./testdata",
	})

	assert.Equal(t, "./testdata", config.GetOptions().GetEnvPath())
	assert.Equal(t, "gp", config.GetOptions().GetPrefix())
}

func TestConfig_ServiceProvider(t *testing.T) {
	app := foundation.NewApplication("./")

	app.Register(NewConfigProvider(app))

	app.Make("config").(*Config).Add("test", "test")
	fmt.Println(app.Make("config").(*Config).Get("test"))

	// assert.Equal(t, )
	fmt.Println(app)
}

func TestConfig_Facades(t *testing.T) {
	app := foundation.NewApplication("./")

	app.Register(NewConfigProvider(app))

	Facade().Add("test", "test")

	assert.Equal(t, "test", Facade().Get("test"))
}
