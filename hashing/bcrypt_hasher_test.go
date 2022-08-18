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

	hash2 := &BcryptHasher{}
	hashedValue2 := hash2.MustMake(value)

	assert.Equal(t, true, hash2.Check(value, hashedValue2))
}
