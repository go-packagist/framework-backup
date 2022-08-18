package redis

type Manager struct {
	config      *Config
	connections map[string]Connection
}

type Config struct {
	Default     string
	Connections map[string]map[string]interface{}
}

func NewManager(c *Config) *Manager {
	return &Manager{
		config:      c,
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
	cfg := m.config.Connections[name]

	if cfg == nil {
		panic("redis connection not found: " + name)
	}

	// driver
	switch cfg["driver"].(string) {
	case "redis":
		return m.createRedisConnection(cfg)
	default:
		panic("redis driver not found: " + cfg["driver"].(string))
	}
}

func (m *Manager) createRedisConnection(config map[string]interface{}) Connection {
	conn := NewRedisConnection()

	conn.Connect(config)

	return conn
}

func (m *Manager) getDefaultName() string {
	return m.config.Default
}
