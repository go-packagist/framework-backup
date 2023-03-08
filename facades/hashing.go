package facades

import (
	"fmt"
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
		return nil, fmt.Errorf("[hash] is not a hashing manager")
	}
}

// MustHash returns the hashing manager.
func MustHash() *hashing.Manager {
	hash, err := Hash()

	if err != nil {
		panic(err)
	}

	return hash
}
