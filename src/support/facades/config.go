package facades

import (
	"github.com/go-packagist/config"
	"github.com/go-packagist/foundation"
)

func Config() *config.Config {
	return foundation.App().Make("config").(*config.Config)
}
