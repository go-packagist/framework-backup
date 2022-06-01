package facades

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/hashing"
)

// Hash returns the hashing manager.
func Hash() (*hashing.HashManager, error) {
	hash, err := foundation.App().Make("hash")
	if err != nil {
		return nil, err
	}

	return hash.(*hashing.HashManager), nil
}

// MustHash returns the hashing manager.
func MustHash() *hashing.HashManager {
	hash, _ := Hash()

	return hash
}

// BcryptHasher returns the bcrypt hasher.
func BcryptHasher() (*hashing.BcryptHasher, error) {
	hash, err := foundation.App().Make("hash.bcrypt")
	if err != nil {
		return nil, err
	}

	return hash.(*hashing.BcryptHasher), nil
}

// MustBcryptHasher returns the bcrypt hasher.
func MustBcryptHasher() *hashing.BcryptHasher {
	hash, _ := BcryptHasher()

	return hash
}
