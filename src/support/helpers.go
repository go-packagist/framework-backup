package support

import (
	"time"
)

// With Return the given value, optionally passed through the given callback.
// If the `fn` is nil, the value is returned unmodified.
// If the `fn` is not nil, the value is passed through the callback, and the result is returned.
//
// Example:

// 	With(1, func(i interface{}) interface{} {
// 		return i.(int) + 2
// 	}, func(i interface{}) interface{} {
// 		return i.(int) * 100
// 	})
//
func With(value interface{}, fn ...func(interface{}) interface{}) interface{} {
	for _, f := range fn {
		value = f(value)
	}

	return value
}

// Default Return the given value, or the given default value if the given value is nil.
//
// Example:
//
//	Default(nil, 1)
// 	Default(1, 1)
//
func Default(value interface{}, defaultValue interface{}) interface{} {
	if value == nil {
		return defaultValue
	}

	return value
}

// Retry Retry an operation a given number of times
//
// Example:
//
// 	Retry(3, func() error {
// 		return nil
// 	})
//
func Retry(times int, fn func() error, sleepTime ...time.Duration) error {
	for i := 0; i < times; i++ {
		err := fn()

		if err == nil {
			return nil
		}

		if i == times-1 {
			return err
		}

		if len(sleepTime) > 0 {
			time.Sleep(sleepTime[0])
		}
	}

	return fn() // 实际不会执行到这里
}
