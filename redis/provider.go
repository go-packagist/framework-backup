package redis

import (
	"github.com/go-packagist/framework/container"
)

type redisProvider struct {
	container *container.Container
}

var _ container.Provider = (*redisProvider)(nil)

func NewRedisProvider(c *container.Container) container.Provider {
	return &redisProvider{
		container: c,
	}
}

func (r *redisProvider) Register() {
	r.container.Singleton("redis", func(c *container.Container) interface{} {
		return NewManager(c)
	})
}
