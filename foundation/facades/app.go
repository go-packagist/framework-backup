package facades

import "github.com/go-packagist/framework/foundation"

// App returns the application instance.
func App() *foundation.Application {
	return foundation.App()
}
