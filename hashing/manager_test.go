package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager_Driver(t *testing.T) {
	m := NewHashManager(map[string]interface{}{
		"driver": "bcrypt",
	})

	// Default Driver
	value := "123456"
	hashedValue := m.Driver().MustMake(value)
	assert.True(t, m.Driver().Check(value, hashedValue))

	// Bcrypt Driver
	hashedValue2 := m.Driver("bcrypt").MustMake(value)
	assert.True(t, m.Driver("bcrypt").Check(value, hashedValue2))

	// Md5 Driver
	hashedValue3 := m.Driver("md5").MustMake(value)
	assert.True(t, m.Driver("md5").Check(value, hashedValue3))
}
