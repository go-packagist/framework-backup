# redis

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/framework)](https://goreportcard.com/report/github.com/go-packagist/framework)
[![tests](https://github.com/go-packagist/framework/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/framework/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/framework/redis)](https://pkg.go.dev/github.com/go-packagist/framework/redis)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

_——The idea came from [Laravel](https://github.com/laravel)_

## Usage

```go
package main

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/redis"
	"github.com/go-packagist/framework/support/facades"
)

func main() {
	app := foundation.NewApplication("./")

	app.Register(config.NewConfigProvider(app))

	facades.MustConfig().Add("redis", map[string]interface{}{
		"default": "redis",
		"connections": map[string]interface{}{
			"redis": map[string]interface{}{
				"driver":   "redis",
				"host":     "localhost",
				"port":     63790,
				"database": 0,
				"password": "",
			},
		},
	})

	app.Register(redis.NewRedisProvider(app))

	facades.MustRedis().Connection().Set("key", "value", time.Second*100)
}
```

