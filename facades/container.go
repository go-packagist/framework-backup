package facades

import "github.com/go-packagist/framework/container"

// Container returns the application container.
func Container() *container.Container {
	return App().Container
}
