package cache

// import (
// 	"github.com/go-packagist/foundation"
// 	"time"
// )
//
// type Store interface {
// 	Get(key string) *Result
// 	Put(key string, value interface{}, expire time.Duration) error
// 	Has(key string) bool
// 	Remember(key string, fc func() interface{}, expire time.Duration) *Result
// 	GC() error
// }
//
// type manager struct {
// 	app    *foundation.Application
// 	caches map[string]Store
// }
//
// func New(app *foundation.Application) *manager {
// 	return &manager{
// 		app:    app,
// 		caches: make(map[string]Store),
// 	}
// }
//
// func (m *manager) Store(name string) Store {
// 	return m.caches[name] = m.get(name)
// }
//
// func (m *manager) get(name string) Store {
// 	if cache, ok := m.caches[name]; ok {
// 		return cache
// 	}
//
// 	return m.caches[name] = m.resolve(name)
// }
//
// func (m *manager) resolve(name string) Store {
// 	config := m.getConfig(name)
//
// }
//
// func (m *manager) getConfig(name string) Config {
// 	return m.app.Config.Get("cache." + name).(Config)
// }
