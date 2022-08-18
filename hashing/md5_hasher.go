package hashing

import (
	"crypto/md5"
	"fmt"
)

type Md5Hasher struct {
}

var _ Hasher = (*Md5Hasher)(nil)

// NewMd5Hasher creates a new md5 hasher instance.
func NewMd5Hasher() *Md5Hasher {
	return &Md5Hasher{}
}

// Make generates a new hashed value.
func (m *Md5Hasher) Make(value string) (string, error) {
	hashedValue := md5.Sum([]byte(value))

	return fmt.Sprintf("%x", hashedValue), nil
}

// MustMake generates a new hashed value.
func (m *Md5Hasher) MustMake(value string) string {
	hashedValue, err := m.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// Check checks the given value and hashed value.
func (m *Md5Hasher) Check(value, hashedValue string) bool {
	return m.MustMake(value) == hashedValue
}
