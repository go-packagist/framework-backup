package container

type Container struct {
	providers map[string]Provider
}

func NewContainer() *Container {
	return &Container{
		providers: make(map[string]Provider),
	}
}

func (c *Container) Register(name string, provider Provider) Provider {
	c.providers[name] = provider

	provider.Register()

	return provider
}

func (c *Container) GetProvider(name string) Provider {
	return c.providers[name]
}

func (c *Container) GetProviders() map[string]Provider {
	return c.providers
}
