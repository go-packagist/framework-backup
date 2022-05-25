package providers

import (
	"github.com/go-packagist/foundation"
	"time"
)

type memoryProvider struct {
	app *foundation.Application
}

var _ foundation.Provider = (*memoryProvider)(nil)

func NewMemoryProvider(app *foundation.Application) foundation.Provider {
	return &memoryProvider{
		app: app,
	}
}

func (m *memoryProvider) Register() {
	m.app.Singleton("memory", func(app *foundation.Application) interface{} {
		return NewMemory(app)
	})
}

type memoryData struct {
	Key    string
	Value  string
	Expire time.Time
}

type Memory struct {
	app   *foundation.Application
	items map[string]memoryData
}

func NewMemory(app *foundation.Application) *Memory {
	m := &Memory{
		app:   app,
		items: make(map[string]memoryData),
	}

	go m.gc()

	return m
}

func (m *Memory) Put(key, value string, expire time.Duration) {
	m.items[key] = memoryData{
		Key:    key,
		Value:  value,
		Expire: time.Now().Add(expire),
	}
}

func (m *Memory) Get(key string) string {
	data, ok := m.items[key]

	if !ok {
		return ""
	}

	return data.Value
}

func (m *Memory) gc() {
	for {
		for key, data := range m.items {
			if time.Now().After(data.Expire) {
				delete(m.items, key)
			}
		}

		time.Sleep(time.Second)
	}
}
