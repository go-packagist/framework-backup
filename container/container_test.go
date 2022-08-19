package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer(t *testing.T) {
	c := NewContainer()

	c.Instance("test", "test")
	assert.Equal(t, "test", c.MustMake("test").(string))
	test, _ := c.Get("test")
	assert.Equal(t, "test", test.(string))

	c.Singleton("test2", func(c *Container) interface{} {
		return "test2"
	})
	assert.Equal(t, "test2", c.MustMake("test2").(string))

	var i int
	c.Singleton("test3", func(c *Container) interface{} {
		i += 1
		return i
	})
	assert.Equal(t, 1, c.MustMake("test3").(int))
	assert.Equal(t, 1, c.MustMake("test3").(int))

	var j int
	c.Bind("test4", func(c *Container) interface{} {
		j += 1
		return j
	}, false)
	assert.Equal(t, 1, c.MustMake("test4").(int))
	assert.Equal(t, 2, c.MustMake("test4").(int))
}
