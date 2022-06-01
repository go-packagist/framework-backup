package config

import (
	"github.com/go-packagist/framework/foundation"
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

	config, err := app.Make("config")
	facade := config.(*Config)

	facade.Add("test", "test")

	assert.Equal(t, "test", facade.Get("test"))
	assert.Nil(t, err)
}

func TestConfig_Add(t *testing.T) {
	app := foundation.NewApplication("./")

	app.Register(NewConfigProvider(app))

	config := app.MustMake("config").(*Config)

	config.Add("app", map[string]interface{}{
		"name":     "test",
		"debug":    true,
		"timezone": "Beijing",
		"map": map[string]interface{}{
			"key": "value",
		},
	})

	assert.Equal(t, "test", config.Get("app.name"))
	assert.Equal(t, true, config.Get("app.debug"))
	assert.Equal(t, "Beijing", config.Get("app.timezone"))
	assert.Equal(t, "value", config.Get("app.map.key"))
	assert.Equal(t, map[string]interface{}{
		"app": map[string]interface{}{
			"name":     "test",
			"debug":    true,
			"timezone": "Beijing",
			"map": map[string]interface{}{
				"key": "value",
			},
		},
	}, config.GetAll())
}
