# foundation

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/foundation)](https://goreportcard.com/report/github.com/go-packagist/foundation)
[![tests](https://github.com/go-packagist/foundation/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/foundation/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/foundation)](https://pkg.go.dev/github.com/go-packagist/foundation)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

_——The idea came from [Laravel](https://github.com/laravel)_

## Installation

```bash
go get github.com/go-packagist/foundation
```

## Usage

- `example/foundation/providers`
    
    ```go
    package providers
    
    import (
        "github.com/go-packagist/foundation"
        "time"
    )
    
    type memoryProvider struct {
        app *foundation.Application
    }
    
    var _ foundation.Provider = (*memoryProvider)(nil)
    
    func NewMemoryProvider(app *foundation.Application) foundation.Provider {
        return &memoryProvider{
            app: app,
        }
    }
    
    func (m *memoryProvider) Register() {
        m.app.Singleton("memory", func(app *foundation.Application) interface{} {
            return NewMemory(app)
        })
    }
    
    type memoryData struct {
        Key    string
        Value  string
        Expire time.Time
    }
    
    type Memory struct {
        app   *foundation.Application
        items map[string]memoryData
    }
    
    func NewMemory(app *foundation.Application) *Memory {
        m := &Memory{
            app:   app,
            items: make(map[string]memoryData),
        }
    
        go m.gc()
    
        return m
    }
    
    func (m *Memory) Put(key, value string, expire time.Duration) {
        m.items[key] = memoryData{
            Key:    key,
            Value:  value,
            Expire: time.Now().Add(expire),
        }
    }
    
    func (m *Memory) Get(key string) string {
        data, ok := m.items[key]
    
        if !ok {
            return ""
        }
    
        return data.Value
    }
    
    func (m *Memory) gc() {
        for {
            for key, data := range m.items {
                if time.Now().After(data.Expire) {
                    delete(m.items, key)
                }
            }
    
            time.Sleep(time.Second)
        }
    }
    ```
  
- `example/foundation/facades`

  ```go
  package facades
  
  import (
      "example/foundation/providers"
      "github.com/go-packagist/foundation"
  )
  
  func Memory() *providers.Memory {
      return foundation.App().Make("memory").(*providers.Memory)
  }
  ```
  
- `main.go`

    ```go
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
        fmt.Println("expire:" + facades.Memory().Get("a")) // using facades
    }
    ```
