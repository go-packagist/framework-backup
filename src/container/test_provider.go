package container

type TestProvider struct {
	Container *Container
}

var _ Provider = (*TestProvider)(nil)

func NewTestProvider(c *Container) *TestProvider {
	return &TestProvider{
		Container: c,
	}
}

func (p *TestProvider) Register() {
	return
}
