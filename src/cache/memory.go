package cache

import (
	"errors"
	"time"
)

type memoryStore struct {
	data map[string]*memoryData
}

type memoryData struct {
	key    string
	value  interface{}
	expire time.Time
}

func NewMemoryStore() Cache {
	store := &memoryStore{
		data: make(map[string]*memoryData),
	}

	store.GC()

	return store
}

func (m *memoryStore) Put(key string, value interface{}, expire time.Duration) error {
	m.data[key] = &memoryData{
		key:    key,
		value:  value,
		expire: time.Now().Add(expire),
	}

	return nil
}

func (m *memoryStore) Get(key string) *Result {
	data, ok := m.data[key]

	if ok {
		return &Result{data.value, nil}
	}

	return &Result{nil, errors.New("key not found")}
}

func (m *memoryStore) Has(key string) bool {
	_, ok := m.data[key]

	return ok
}

func (m *memoryStore) Remember(key string, fc func() interface{}, expire time.Duration) *Result {
	if !m.Has(key) {
		m.Put(key, fc(), expire)
	}

	return m.Get(key)
}

func (m *memoryStore) GC() error {
	go func() {
		for {
			for key, data := range m.data {
				if data.expire.Before(time.Now()) {
					delete(m.data, key)
				}
			}
		}
	}()

	return nil
}
