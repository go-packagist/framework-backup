package config

import "github.com/go-packagist/foundation"

func Facade() *Config {
	return foundation.App().Make("config").(*Config)
}
