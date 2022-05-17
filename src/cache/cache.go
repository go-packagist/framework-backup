package cache

import "time"

type Cache interface {
	Get(key string) *Result
	Put(key string, value interface{}, expire time.Duration) error
	Has(key string) bool
	Remember(key string, fc func() interface{}, expire time.Duration) *Result
	GC() error
}

var stores = make(map[string]Cache)

func Configure(name string, store Cache) {
	stores[name] = store
}

func Store(name string) Cache {
	return stores[name]
}

type (
	MarshalFunc   func(interface{}) ([]byte, error)
	UnmarshalFunc func([]byte, interface{}) error
)
