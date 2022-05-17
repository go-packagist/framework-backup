package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer_Register(t *testing.T) {
	c := NewContainer()

	c.Register("test", NewTestProvider(c))

	assert.Equal(t, map[string]Provider{
		"test": NewTestProvider(c),
	}, c.GetProviders())
}
