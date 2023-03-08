package hashing

type Hasher interface {
	Make(string) (string, error)
	MustMake(string) string
	Check(string, string) bool
}
