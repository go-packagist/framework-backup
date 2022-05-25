package facades

import (
	"example/foundation/providers"
	"github.com/go-packagist/foundation"
)

func Memory() *providers.Memory {
	return foundation.App().Make("memory").(*providers.Memory)
}
