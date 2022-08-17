package facades

import (
	"errors"
	"github.com/go-packagist/framework/config"
)

// Config returns the config facade.
func Config() (*config.Config, error) {
	cfg, err := App().Make("config")

	if err != nil {
		return nil, err
	}

	switch cfg.(type) {
	case *config.Config:
		return cfg.(*config.Config), nil
	default:
		return nil, errors.New("config is not a config")
	}
}

// MustConfig returns the config facade.
func MustConfig() *config.Config {
	cfg, _ := Config()

	return cfg
}
