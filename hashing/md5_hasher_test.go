package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5Hasher(t *testing.T) {
	md5 := NewMd5Hasher()

	value := "123456"
	hashedValue, err := md5.Make(value)

	assert.Nil(t, err)
	assert.True(t, md5.Check(value, hashedValue))
}
