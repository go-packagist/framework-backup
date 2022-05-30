package main

import (
	"fmt"
	"github.com/go-packagist/config"
	"github.com/go-packagist/foundation"
)

func main() {
	app := foundation.NewApplication("./")

	app.Register(config.NewConfigProvider(app))

	fmt.Println(app)
}
