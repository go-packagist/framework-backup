package facades

import "github.com/go-packagist/framework/redis"

// Redis returns the redis manager instance.
func Redis() (*redis.Manager, error) {
	cfg, err := App().Make("redis")

	if err != nil {
		return nil, err
	}

	return cfg.(*redis.Manager), nil
}

// MustRedis returns the redis manager instance.
func MustRedis() *redis.Manager {
	redis, _ := Redis()

	return redis
}
