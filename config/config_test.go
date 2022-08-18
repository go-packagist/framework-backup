package config

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions(t *testing.T) {
	options := &Options{
		EnvPath: "./.testdata",
	}

	assert.Equal(t, "./.testdata", options.GetEnvPath())

	options.Prepare()
	assert.Equal(t, "gp", options.GetPrefix())
}

func createConfig() *Config {
	return New(&Options{
		EnvPath: "./testdata",
	})
}

func TestConfig_Map(t *testing.T) {
	config := createConfig()

	// map
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

func TestConfig_Struct(t *testing.T) {
	config := createConfig()

	type Test struct {
		Name string
	}

	test := &Test{
		Name: "test",
	}

	config.Add("struct", test)

	assert.Equal(t, "test", config.Get("struct").(*Test).Name)
	assert.Equal(t, test, config.Get("struct"))
}

func TestConfig_ServiceProvider(t *testing.T) {
	app := foundation.NewApplication("./")

	app.Register(NewConfigProvider(app.Container))

	config, err := app.Make("config")
	facade := config.(*Config)

	facade.Add("test", "test")

	assert.Equal(t, "test", facade.Get("test"))
	assert.Nil(t, err)
}
