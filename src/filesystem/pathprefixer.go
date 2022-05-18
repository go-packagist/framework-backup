package filesystem

import "strings"

type PathPrefixer struct {
	prefix    string
	separator string
}

// separator is the separator used to separate the prefix from the path.
const separator = "/"

// NewPathPrefixer creates a new PathPrefixer.
func NewPathPrefixer(prefix string) *PathPrefixer {
	prefix = strings.TrimRight(prefix, "\\/")

	return &PathPrefixer{
		prefix:    prefix + separator,
		separator: separator,
	}
}

// PrefixPath Returns a prefixed directory
func (p *PathPrefixer) PrefixPath(path string) string {
	return p.prefix + strings.TrimLeft(path, "\\/")
}
