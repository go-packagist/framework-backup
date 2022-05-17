package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	Configure("memory", NewMemoryStore())
	Store("memory").Put("a", "aaa", time.Second*10)

	assert.Equal(t, "aaa", Store("memory").Get("a").Val)
}
