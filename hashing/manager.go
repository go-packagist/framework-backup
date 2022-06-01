package hashing

import (
	"fmt"
)

type HashManager struct {
	config  map[string]interface{}
	drivers map[string]Hasher
}

// NewManager creates a new hashing manager instance.
// config example:
// 	config := map[string]interface{}{
// 		"driver": "bcrypt",
// 	}
func NewManager(config map[string]interface{}) *HashManager {
	return &HashManager{
		config:  config,
		drivers: make(map[string]Hasher),
	}
}

// Driver gets the hasher instance by driver name.
func (m *HashManager) Driver(driver ...string) Hasher {
	if len(driver) > 0 {
		return m.resolve(driver[0])
	}

	return m.resolve(m.getDefaultDriver())
}

// resolve gets the hasher instance by name.
func (m *HashManager) resolve(driver string) Hasher {
	hasher, ok := m.drivers[driver]

	if ok {
		return hasher
	}

	switch driver {
	case "bcrypt":
		hasher = m.createBcryptHasher()
	default:
		panic(fmt.Sprintf("hashing driver %s is not supported", driver))
	}

	m.drivers[driver] = hasher

	return hasher
}

// createBcryptHasher creates a new bcrypt hasher instance.
func (m *HashManager) createBcryptHasher() Hasher {
	return NewBcryptHasher()
}

// getDefaultDriver gets the default driver name.
func (m *HashManager) getDefaultDriver() string {
	return m.config["driver"].(string)
}
