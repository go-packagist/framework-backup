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

func TestConfig_Add(t *testing.T) {
	app := foundation.NewApplication("./")

	app.Register(NewConfigProvider(app))

	Facade().Add("app", map[string]interface{}{
		"name":     "test",
		"debug":    true,
		"timezone": "Beijing",
	})

	assert.Equal(t, "test", Facade().Get("app.name"))
	assert.Equal(t, true, Facade().Get("app.debug"))
	assert.Equal(t, "Beijing", Facade().Get("app.timezone"))
	assert.Equal(t, map[string]interface{}{
		"app": map[string]interface{}{
			"name":     "test",
			"debug":    true,
			"timezone": "Beijing",
		},
	}, Facade().GetAll())
}
