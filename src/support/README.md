# support

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/support)](https://goreportcard.com/report/github.com/go-packagist/support)
[![tests](https://github.com/go-packagist/support/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/support/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/support)](https://pkg.go.dev/github.com/go-packagist/support)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

_——The idea came from [Laravel](https://github.com/laravel)_

## Installation

```bash
go get github.com/go-packagist/filesystem
```

## Usage

```go
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
```