package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore(client *redis.Client) Cache {
	return &redisStore{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *redisStore) Get(key string) *Result {
	val, err := r.client.Get(r.ctx, key).Result()

	// fixed: val is empty string when key is not exist
	if err != nil {
		return &Result{nil, err}
	}

	return &Result{val, err}
}

func (r *redisStore) Put(key string, value interface{}, expire time.Duration) error {
	return r.client.SetEX(r.ctx, key, value, expire).Err()
}

func (r *redisStore) Has(key string) bool {
	return r.client.Exists(r.ctx, key).Val() == 1
}

func (r *redisStore) Remember(key string, fc func() interface{}, expire time.Duration) *Result {
	if !r.Has(key) {
		val := fc()

		r.Put(key, val, expire)

		return &Result{val, nil}
	}

	return r.Get(key)
}

func (r *redisStore) GC() error {
	return nil
}
