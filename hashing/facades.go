package hashing

import (
	"github.com/go-packagist/framework/foundation"
)

// FacadeHash returns the hashing manager.
func FacadeHash() (*HashManager, error) {
	hash, err := foundation.App().Make("hash")
	if err != nil {
		return nil, err
	}

	return hash.(*HashManager), nil
}

// FacadeMustHash returns the hashing manager.
func FacadeMustHash() *HashManager {
	hash, _ := FacadeHash()

	return hash
}

// FacedeHashBcrypt returns the bcrypt hasher.
func FacedeHashBcrypt() (*BcryptHasher, error) {
	hash, err := foundation.App().Make("hash.bcrypt")
	if err != nil {
		return nil, err
	}

	return hash.(*BcryptHasher), nil
}

// FacadeMustHashBcrypt returns the bcrypt hasher.
func FacadeMustHashBcrypt() *BcryptHasher {
	hash, _ := FacedeHashBcrypt()

	return hash
}
