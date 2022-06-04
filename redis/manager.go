package redis

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
)

type Manager struct {
	app         *foundation.Application
	connections map[string]Connection
}

func NewManager(app *foundation.Application) *Manager {
	return &Manager{
		app:         app,
		connections: make(map[string]Connection),
	}
}

func (m *Manager) Connection(name ...string) Connection {
	if len(name) == 0 {
		name = []string{m.getDefaultName()}
	}

	if conn, ok := m.connections[name[0]]; ok {
		return conn
	}

	m.connections[name[0]] = m.resolve(name[0])

	return m.connections[name[0]]
}

func (m *Manager) resolve(name string) Connection {
	config := m.app.MustMake("config").(*config.Config).
		Get("redis.connections." + name).(map[string]interface{})

	if config == nil {
		panic("redis connection not found: " + name)
	}

	// driver
	switch config["driver"].(string) {
	case "redis":
		return m.createRedisConnection(config)
	default:
		panic("redis driver not found: " + config["driver"].(string))
	}
}

func (m *Manager) createRedisConnection(config map[string]interface{}) Connection {
	conn := NewRedisConnection()

	conn.Connect(config)

	return conn
}

func (m *Manager) getDefaultName() string {
	return m.app.MustMake("config").(*config.Config).Get("redis.default").(string)
}
