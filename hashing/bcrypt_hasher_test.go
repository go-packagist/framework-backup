package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBcryptHasher(t *testing.T) {
	value := "123456"

	hash := NewBcryptHasher()
	hashedValue, _ := hash.Make(value)

	assert.Equal(t, true, hash.Check(value, hashedValue))
}
