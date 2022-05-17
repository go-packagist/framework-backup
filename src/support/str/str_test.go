package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStr_Contains(t *testing.T) {
	assert.True(t, Contains("abc", "a"))
	assert.True(t, Contains("abc", "abc"))
	assert.True(t, Contains("abc", ""))
	assert.False(t, Contains("abc", "abcd"))
}
