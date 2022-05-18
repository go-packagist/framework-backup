package coroutine

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	p := NewParallel(0)

	p.Add(func() interface{} {
		time.Sleep(2 * time.Second)
		return 2
	})

	p.Add(func() interface{} {
		time.Sleep(2 * time.Second)
		return 1
	})

	p.Add(func() interface{} {
		time.Sleep(1 * time.Second)

		return 3
	})

	p.Add(func() interface{} {
		return 4
	})

	assert.Equal(t, []interface{}{2, 1, 3, 4}, p.Wait())
}

func TestParallel_Limit(t *testing.T) {

}
