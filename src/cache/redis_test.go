package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createRedisStore() Cache {
	return NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "localhost:63790",
	}))
}

func TestRedisStore_PutAndGet(t *testing.T) {
	redis := createRedisStore()

	redis.Put("a", "aaa", time.Second*1)
	assert.Equal(t, "aaa", redis.Get("a").Value())

	redis.Put("b", "bbb", time.Second*1)
	time.Sleep(time.Second * 2)
	assert.Equal(t, nil, redis.Get("b").Value())
}

func TestRedisStore_Has(t *testing.T) {
	redis := createRedisStore()

	redis.Put("has", "has", time.Second*1)
	assert.True(t, redis.Has("has"))
	assert.False(t, redis.Has("has2"))
}

func TestRedisStore_Remember(t *testing.T) {
	redis := createRedisStore()

	assert.Equal(t, nil, redis.Get("remember").Value())

	redis.Remember("remember", func() interface{} {
		return "remember"
	}, time.Second*10)
	assert.Equal(t, "remember", redis.Get("remember").Value())

	redis.Remember("remember", func() interface{} {
		return "remember2"
	}, time.Second*10)
	assert.Equal(t, "remember", redis.Get("remember").Value())
}
