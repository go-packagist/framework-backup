package facades

import (
	"github.com/go-packagist/framework/config"
)

// Config returns the config facade.
func Config() (*config.Config, error) {
	cfg, err := App().Make("config")

	if err != nil {
		return nil, err
	}

	return cfg.(*config.Config), nil
}

// MustConfig returns the config facade.
func MustConfig() *config.Config {
	config, _ := Config()

	return config
}
