package container

import (
	"errors"
	"fmt"
	"sync"
)

type ConcreteFunc func(*Container) interface{}

type binding struct {
	abstract string
	concrete ConcreteFunc
	shared   bool
}

type Container struct {
	bindings  map[string]binding
	instances map[string]interface{}

	rwlock *sync.RWMutex
}

func NewContainer() *Container {
	c := &Container{
		bindings:  make(map[string]binding),
		instances: make(map[string]interface{}),
		rwlock:    &sync.RWMutex{},
	}

	return c
}

// Singleton Register a shared binding in the container.
func (c *Container) Singleton(abstract string, concrete ConcreteFunc) {
	c.Bind(abstract, concrete, true)
}

// Bind Register a binding with the container.
func (c *Container) Bind(abstract string, concrete ConcreteFunc, shared bool) {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()

	c.bindings[abstract] = binding{
		abstract: abstract,
		concrete: concrete,
		shared:   shared,
	}
}

// Instance Set the given type to the container.
func (c *Container) Instance(abstract string, concrete interface{}) {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()

	c.instances[abstract] = concrete
}

// Make Resolve the given type from the container.
func (c *Container) Make(abstract string) (interface{}, error) {
	return c.Resolve(abstract)
}

// MustMake Resolve the given type from the container or panic.
func (c *Container) MustMake(abstract string) interface{} {
	concrete, _ := c.Make(abstract)

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
