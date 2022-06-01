package hashing

type Hasher interface {
	Make(value string) (string, error)
	MustMake(value string) string
	Check(value, hashedValue string) bool
}
