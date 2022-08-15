package foundation

import "github.com/go-packagist/framework/container"

type TestProvider struct {
	app *Application
}

var _ Provider = (*TestProvider)(nil)

func NewTestProvider(app *Application) Provider {
	return &TestProvider{
		app: app,
	}
}

func (p *TestProvider) Register() {
	p.app.Singleton("test", func(c *container.Container) interface{} {
		return NewTestService(c)
	})

	p.app.Bind("test2", func(c *container.Container) interface{} {
		return NewTestService(c)
	}, false)
}

// TestService is a test service
type TestService struct {
	c       *container.Container
	content string
}

func NewTestService(c *container.Container) *TestService {
	return &TestService{
		c:       c,
		content: "",
	}
}

func (s *TestService) Container() *container.Container {
	return s.c
}

func (s *TestService) WriteContent(content string) *TestService {
	s.content = content

	return s
}

func (s *TestService) ReadContent() string {
	return s.content
}
