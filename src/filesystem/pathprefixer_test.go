package filesystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathPrefixer(t *testing.T) {
	prefixer := NewPathPrefixer("/foo/bar/")
	prefixer2 := NewPathPrefixer("/foo/bar")

	assert.Equal(t, "/foo/bar/baz", prefixer.PrefixPath("/baz"))
	assert.Equal(t, "/foo/bar/baz/", prefixer.PrefixPath("/baz/"))

	assert.Equal(t, "/foo/bar/baz", prefixer2.PrefixPath("/baz"))
	assert.Equal(t, "/foo/bar/baz/", prefixer2.PrefixPath("/baz/"))
}
