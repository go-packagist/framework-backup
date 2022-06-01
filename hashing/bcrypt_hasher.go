package hashing

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct {
}

var _ Hasher = (*BcryptHasher)(nil)

// NewBcryptHasher returns a new bcrypt hasher.
func NewBcryptHasher() Hasher {
	return &BcryptHasher{}
}

// Make returns the hashed value.
func (b *BcryptHasher) Make(value string) (string, error) {
	return b.MakeWithCost(value, bcrypt.DefaultCost)
}

// MustMake returns the hashed value.
func (b *BcryptHasher) MustMake(value string) string {
	hashedValue, err := b.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// MakeWithCost returns the hashed value with the given cost.
func (b *BcryptHasher) MakeWithCost(value string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), cost)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// Check returns true if the value matches the hashed value.
func (b *BcryptHasher) Check(value, hashedValue string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value)) == nil
}
