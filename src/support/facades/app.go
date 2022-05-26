package facades

import "github.com/go-packagist/foundation"

// App returns the application instance.
func App() *foundation.Application {
	return foundation.App()
}
