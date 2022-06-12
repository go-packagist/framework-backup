package database

import (
	"fmt"
	"sync"
)

type Manager struct {
	config     *Config
	connectors map[string]Connector
	locker     *sync.RWMutex
}

type Config struct {
	Default     string
	Connections map[string]interface{}
}

func NewManager(config *Config) *Manager {
	return &Manager{
		config:     config,
		connectors: make(map[string]Connector),
		locker:     &sync.RWMutex{},
	}
}

func (m *Manager) Connection(name ...string) Connector {
	if len(name) > 0 {
		return m.resolve(name[0])
	}

	return m.resolve(m.getDefaultName())
}

// resolve gets the hasher instance by name.
func (m *Manager) resolve(name string) Connector {
	m.locker.Lock()
	defer m.locker.Unlock()

	connector, ok := m.connectors[name]

	if ok {
		return connector
	}

	config := m.getConfig(name)

	switch config.(type) {
	case *MySqlConfig:
		connector = m.createMysqlConnector(config.(*MySqlConfig))
	default:
		panic("driver is not supported")
	}

	fmt.Println(connector)

	m.connectors[name] = connector

	return connector
}

// createMysqlConnector creates a new mysql connector instance.
func (m *Manager) createMysqlConnector(config *MySqlConfig) Connector {
	return NewMySQLConnector(config)
}

// getDefaultName gets the default driver name.
func (m *Manager) getDefaultName() string {
	return m.config.Default
}

// getConfig gets the config by name.
func (m *Manager) getConfig(name string) interface{} {
	return m.config.Connections[name]
}
