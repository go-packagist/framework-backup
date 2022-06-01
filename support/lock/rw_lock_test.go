package lock

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRWLock_Run(t *testing.T) {
	l := NewRwLock()
	d := make(map[string]string)

	for i := 0; i < 100; i++ {
		go func(i int) {
			// d["test"+strconv.Itoa(i)] = "test" + strconv.Itoa(i)
			l.Run(func() {
				d["test"+strconv.Itoa(i)] = "test" + strconv.Itoa(i)
			})
		}(i)
	}

	for {
		if len(d) == 100 {
			break
		}
	}

	assert.Equal(t, 100, len(d))
}
