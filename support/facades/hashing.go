package facades

import (
	"github.com/go-packagist/framework/hashing"
)

// Hash returns the hashing manager.
func Hash() (*hashing.HashManager, error) {
	hash, err := App().Make("hash")
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
	hash, err := App().Make("hasher.bcrypt")
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

// Md5Hasher returns the md5 hasher.
func Md5Hasher() (*hashing.Md5Hasher, error) {
	hash, err := App().Make("hasher.md5")
	if err != nil {
		return nil, err
	}

	return hash.(*hashing.Md5Hasher), nil
}

// MustMd5Hasher returns the md5 hasher.
func MustMd5Hasher() *hashing.Md5Hasher {
	hash, _ := Md5Hasher()

	return hash
}
