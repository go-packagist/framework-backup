# hashing

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/framework)](https://goreportcard.com/report/github.com/go-packagist/framework)
[![tests](https://github.com/go-packagist/framework/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/framework/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/framework/hashing)](https://pkg.go.dev/github.com/go-packagist/framework/hashing)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

_——The idea came from [Laravel](https://github.com/laravel)_

## Usage

```go
package main

import (
	"github.com/go-packagist/framework/hashing"
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/support/facades"
)

func main() {
	// use hashing.NewBcryptHasher
	bcrypt := hashing.NewBcryptHasher()

	bcrypt.Make("password")
	bcrypt.MustMake("password")
	bcrypt.Check("password", "hashed password")

	// use Manager
	m := hashing.NewManager(map[string]interface{}{
		"driver": "bcrypt",
	})

	m.Driver("bcrypt").Make("password")
	m.Driver().Make("password")

	// use facades
	app := foundation.NewApplication("./")

	app.Register(config.NewConfigProvider(app))
	app.MustMake("config").(*config.Config).Set("hashing", map[string]interface{}{
		"driver": "bcrypt",
	})

	app.Register(hashing.NewHashProvider(app))

	facades.MustHash().Driver().Make("password")
	facades.MustHash().Driver("bcrypt").Make("password")
	facades.MustHash().Driver("md5").Make("password")
	facades.BcryptHasher().Make("password")
	facades.Md5Hasher().Make("password")
	app.MustMake("hash").(*hashing.HashManager).Driver("bcrypt").Make("password")
}
```

