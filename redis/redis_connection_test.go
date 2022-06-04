package redis

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createRedisConnection() Connection {
	app := foundation.NewApplication("./")

	// Register Config
	app.Register(config.NewConfigProvider(app))

	// Configure Redis Connection
	app.MustMake("config").(*config.Config).Add("redis", map[string]interface{}{
		"default": "redis",
		"connections": map[string]interface{}{
			"redis": map[string]interface{}{
				"driver":   "redis",
				"host":     "localhost",
				"port":     63790,
				"password": "",
				"database": 0,
			},
		},
	})

	// Redis Connection
	redis := NewRedisConnection()
	redis.Connect(app.MustMake("config").(*config.Config).
		Get("redis.connections.redis").(map[string]interface{}))

	return redis
}

func TestRedisConnection_Connect(t *testing.T) {
	redis := createRedisConnection()

	// cmd
	echo, _ := redis.Echo("Hello world") // echo
	ping, _ := redis.Ping()              // ping

	// assert
	assert.Equal(t, "Hello world", echo)
	assert.Equal(t, "PONG", ping)
}

func TestRedisConnection_Base(t *testing.T) {
	redis := createRedisConnection()

	// cmd
	set, _ := redis.Set("test", "test", time.Second*10)
	set2, _ := redis.Set("test2", "test2", time.Second*10)
	exists, _ := redis.Exists("test", "test2", "test3")
	del, _ := redis.Del("test", "test2")
	del2, _ := redis.Del("test3")
	exists2, _ := redis.Exists("test", "test2", "test3")

	// assert
	assert.Equal(t, "OK", set)
	assert.Equal(t, "OK", set2)
	assert.Equal(t, int64(2), exists)
	assert.Equal(t, int64(2), del)
	assert.Equal(t, int64(0), del2)
	assert.Equal(t, int64(0), exists2)
}
