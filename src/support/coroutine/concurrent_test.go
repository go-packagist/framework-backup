package coroutine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	c := NewConcurrent(10)

	// Concurrently info
	assert.Equal(t, 10, c.Capacity())
	assert.Equal(t, 0, c.Length())
	assert.False(t, c.IsFull())
	assert.True(t, c.IsEmpty())

	// Concurrently Create
	var total int

	go func() {
		time.Sleep(time.Second)

		assert.Equal(t, 10, c.Length())
		assert.Equal(t, 0, total)
		assert.True(t, c.IsFull())

		time.Sleep(1500 * time.Millisecond)
		assert.Equal(t, 5, c.Length())
		assert.Equal(t, 10, total)

		fmt.Println(c.Length(), total, c.IsFull(), c.IsEmpty())
	}()

	for i := 0; i < 15; i++ {
		c.Create(func() {
			time.Sleep(time.Second * 2)

			total++
			// fmt.Println(total, time.Now())
		})
	}

	for {
		if c.IsEmpty() {
			break
		}
	}
}
