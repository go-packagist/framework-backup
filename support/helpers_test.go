package support

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHelpers_With(t *testing.T) {
	r1 := With(1, func(i interface{}) interface{} {
		return i.(int) + 2
	}, func(i interface{}) interface{} {
		return i.(int) * 100
	})

	r2 := With("222", func(i interface{}) interface{} {
		return i.(string) + "333"
	}, func(i interface{}) interface{} {
		return i.(string) + "444"
	})

	r3 := With(1)

	assert.Equal(t, 300, r1)
	assert.Equal(t, "222333444", r2)
	assert.Equal(t, 1, r3)
}

func TestHelpers_Default(t *testing.T) {
	r1 := Default(nil, 1)
	r2 := Default(nil, "222")
	r3 := Default(r2, 1)

	assert.Equal(t, 1, r1)
	assert.Equal(t, "222", r2)
	assert.Equal(t, "222", r3)
}

func TestHelpers_Retry(t *testing.T) {
	var i, ii int

	// retry 3 times and succeed
	Retry(3, func() error {
		i++

		if i <= 1 {
			return errors.New("error")
		}

		return nil
	})

	// retry 3 times and fail
	Retry(3, func() error {
		ii++

		if ii <= 3 {
			return errors.New("error")
		}

		return nil
	})

	// retry 3 times and sleep
	t1 := time.Now()
	Retry(3, func() error {
		return errors.New("error")
	}, time.Second*2)
	t2 := time.Now()

	assert.Equal(t, 2, i)
	assert.Equal(t, 3, ii)

	// 无法准确获取时间差，只能取区间
	assert.True(t, t2.Sub(t1) > time.Second*3)
	assert.True(t, t2.Sub(t1) < time.Second*5)
}
