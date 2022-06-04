package redis

import "github.com/go-packagist/framework/foundation"

type redisProvider struct {
	app *foundation.Application
}

var _ foundation.Provider = (*redisProvider)(nil)

func NewRedisProvider(app *foundation.Application) foundation.Provider {
	return &redisProvider{
		app: app,
	}
}

func (r *redisProvider) Register() {
	r.app.Singleton("redis", func(app *foundation.Application) interface{} {
		return NewManager(app)
	})
}
