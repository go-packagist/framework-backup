package container

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type ConcreteFunc func(*Container) interface{}

type binding struct {
	abstract string
	concrete ConcreteFunc
	shared   bool
}

type Container struct {
	providers []Provider
	bindings  map[string]binding
	instances map[string]interface{}

	locker *sync.RWMutex
}

func NewContainer() *Container {
	c := &Container{
		providers: []Provider{},
		bindings:  make(map[string]binding),
		instances: make(map[string]interface{}),
		locker:    &sync.RWMutex{},
	}

	return c
}

// Register registers a provider with the application.
func (c *Container) Register(provider Provider) {
	if c.providerIsRegistered(provider) {
		return
	}

	provider.Register()

	c.providerMarkAsRegistered(provider)

	// todo
	// 1. bind
	// 2. boot
}

// providerIsRegistered return provider is registered
func (c *Container) providerIsRegistered(provider Provider) bool {
	for _, providerRegistered := range c.providers {
		if reflect.DeepEqual(providerRegistered, provider) {
			return true
		}
	}

	return false
}

// providerMarkAsRegistered provider mark as registered.
func (c *Container) providerMarkAsRegistered(provider Provider) {
	c.providers = append(c.providers, provider)
}

// GetProviders returns all registered providers.
func (c *Container) GetProviders() []Provider {
	return c.providers
}

// Singleton Register a shared binding in the container.
func (c *Container) Singleton(abstract string, concrete ConcreteFunc) {
	c.Bind(abstract, concrete, true)
}

// Bind Register a binding with the container.
func (c *Container) Bind(abstract string, concrete ConcreteFunc, shared bool) {
	c.locker.Lock()
	defer c.locker.Unlock()

	c.bindings[abstract] = binding{
		abstract: abstract,
		concrete: concrete,
		shared:   shared,
	}
}

// Instance Set the given type to the container.
func (c *Container) Instance(abstract string, concrete interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()

	c.instances[abstract] = concrete
}

// Get returns the instance of the given type from the container.
func (c *Container) Get(abstract string) (interface{}, error) {
	return c.Resolve(abstract)
}

// Make Resolve the given type from the container.
func (c *Container) Make(abstract string) (interface{}, error) {
	return c.Resolve(abstract)
}

// MustMake Resolve the given type from the container or panic.
func (c *Container) MustMake(abstract string) interface{} {
	concrete, err := c.Make(abstract)

	if err != nil {
		panic(err)
	}

	return concrete
}

// Resolve the given type from the container.
func (c *Container) Resolve(abstract string) (interface{}, error) {
	// instance
	instance, ok := c.instances[abstract]
	if ok {
		return instance, nil
	}

	// binding
	binding, ok2 := c.bindings[abstract]
	if !ok2 {
		return nil, errors.New(fmt.Sprintf("[%s] binding not found", abstract))
	}

	// concrete(app)
	concrete := binding.concrete(c)

	if c.isShared(abstract) {
		c.Instance(abstract, concrete)
	}

	return concrete, nil
}

// isShared Determine if a given type is shared.
func (c *Container) isShared(abstract string) bool {
	return c.bindings[abstract].shared
}
