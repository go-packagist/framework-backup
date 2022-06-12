# database

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/framework)](https://goreportcard.com/report/github.com/go-packagist/framework)
[![tests](https://github.com/go-packagist/framework/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/framework/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/framework/database)](https://pkg.go.dev/github.com/go-packagist/framework/database)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

_——The idea came from [Laravel](https://github.com/laravel)_

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-packagist/framework/database"
	"gorm.io/gorm"
)

func main() {
	m := database.NewManager(&database.Config{
		Default: "mysql",
		Connections: map[string]interface{}{
			"mysql": &database.MySqlConfig{
				Host:     "localhost",
				Port:     3306,
				Username: "root",
				Password: "123456",
				Database: "test",
			},
		},
	})

	type User struct {
		ID       int    `gorm:"primary_key"`
		Username string `gorm:"column:username"`
	}

	var user User
	m.Connection().DB().(*gorm.DB).First(&user)

	fmt.Println(user)
}
```

