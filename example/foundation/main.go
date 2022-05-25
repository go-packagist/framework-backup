package main

import (
	"example/foundation/facades"
	"example/foundation/providers"
	"fmt"
	"github.com/go-packagist/foundation"
	"time"
)

func main() {
	foundation.NewApplication("./")

	foundation.App().Register(providers.NewMemoryProvider(foundation.App()))

	foundation.App().Make("memory").(*providers.Memory).Put("a", "111", time.Second*2)
	fmt.Println("put:" + foundation.App().Make("memory").(*providers.Memory).Get("a"))

	time.Sleep(time.Second * 3)
	fmt.Println("expire:" + facades.Memory().Get("a"))
}
