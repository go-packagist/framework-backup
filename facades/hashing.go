package facades

import (
	"errors"
	"github.com/go-packagist/framework/hashing"
)

// Hash returns the hashing manager.
func Hash() (*hashing.Manager, error) {
	hash, err := App().Make("hash")
	if err != nil {
		return nil, err
	}

	switch hash.(type) {
	case *hashing.Manager:
		return hash.(*hashing.Manager), nil
	default:
		return nil, errors.New("hash is not a hash manager")
	}
}

// MustHash returns the hashing manager.
func MustHash() *hashing.Manager {
	hash, _ := Hash()

	return hash
}

// BcryptHasher returns the bcrypt hasher.
func BcryptHasher() (*hashing.BcryptHasher, error) {
	hash, err := App().Make("hasher.bcrypt")
	if err != nil {
		return nil, err
	}

	switch hash.(type) {
	case *hashing.BcryptHasher:
		return hash.(*hashing.BcryptHasher), nil
	default:
		return nil, errors.New("hash is not a bcrypt hasher")
	}
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

	switch hash.(type) {
	case *hashing.Md5Hasher:
		return hash.(*hashing.Md5Hasher), nil
	default:
		return nil, errors.New("hash is not a md5 hasher")
	}
}

// MustMd5Hasher returns the md5 hasher.
func MustMd5Hasher() *hashing.Md5Hasher {
	hash, _ := Md5Hasher()

	return hash
}
