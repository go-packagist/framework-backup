# filesystem

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/filesystem)](https://goreportcard.com/report/github.com/go-packagist/filesystem)
[![tests](https://github.com/go-packagist/filesystem/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/filesystem/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/filesystem)](https://pkg.go.dev/github.com/go-packagist/filesystem)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

_——The idea came from [Laravel](https://github.com/laravel)_

> 未完结版

## Installation

```bash
go get github.com/go-packagist/filesystem
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-packagist/filesystem"
)

func main() {
	fs := filesystem.NewManager(&filesystem.Config{
		Default: "local",
		Disk: map[string]interface{}{
			"local": &filesystem.LocalDriveConfig{
				Root: "temp",
			},
		},
	})

	fmt.Println(fs.Disk("local").Exists("test.txt"))
}

```
