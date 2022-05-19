package main

import (
	"fmt"
	"github.com/go-packagist/support/coroutine"
	"time"
)

func main() {
	c := coroutine.NewConcurrent(10)

	for i := 1; i <= 100; i++ {
		ii := i
		c.Create(func() {
			time.Sleep(time.Second)

			fmt.Println(ii, time.Now())
		})
	}

	for {
		if c.IsEmpty() {
			break
		}
	}
}
