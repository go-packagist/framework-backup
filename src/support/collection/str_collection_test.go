package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrCollection_Base(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	// Len, Size, Count
	assert.Equal(t, 3, c.Len())
	assert.Equal(t, 3, c.Size())
	assert.Equal(t, 3, c.Count())

	// Add, Item
	c.Add("d")
	assert.Equal(t, []string{
		"a", "b", "c", "d",
	}, c.Items())

	// Remove
	c.Remove("b")
	assert.Equal(t, []string{
		"a", "c", "d",
	}, c.Items())
	c.Add("b").Add("b")
	assert.Equal(t, []string{
		"a", "c", "d", "b", "b",
	}, c.Items())
	c.Remove("b")
	assert.Equal(t, []string{
		"a", "c", "d", "b",
	}, c.Items())
	c.Add("b")
	c.RemoveAll("b")
	assert.Equal(t, []string{
		"a", "c", "d",
	}, c.Items())
}
