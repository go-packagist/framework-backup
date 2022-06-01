package config

import "github.com/go-packagist/framework/foundation"

// Facade returns the config facade.
func Facade() (*Config, error) {
	config, err := foundation.App().Make("config")

	if err != nil {
		return nil, err
	}

	return config.(*Config), nil
}

// MustFacade returns the config facade.
func MustFacade() *Config {
	config, err := Facade()

	if err != nil {
		panic(err)
	}

	return config
}
