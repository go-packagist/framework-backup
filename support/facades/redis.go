package facades

import (
	"errors"
	"github.com/go-packagist/framework/redis"
)

// Redis returns the redis manager instance.
func Redis() (*redis.Manager, error) {
	rds, err := App().Make("redis")

	if err != nil {
		return nil, err
	}

	switch rds.(type) {
	case *redis.Manager:
		return rds.(*redis.Manager), nil
	default:
		return nil, errors.New("rds is not a redis manager")
	}
}

// MustRedis returns the redis manager instance.
func MustRedis() *redis.Manager {
	rds, _ := Redis()

	return rds
}
