package foundation

type TestProvider struct {
	app *Application
}

var _ Provider = (*TestProvider)(nil)

func NewTestProvider(app *Application) *TestProvider {
	return &TestProvider{
		app: app,
	}
}

func (p *TestProvider) Register() {
	p.app.services["test"] = NewTestService(p.app)
}

type TestService struct {
	app     *Application
	content string
}

func NewTestService(app *Application) *TestService {
	return &TestService{
		app:     app,
		content: "",
	}
}

func (s *TestService) Application() *Application {
	return s.app
}

func (s *TestService) WriteContent(content string) *TestService {
	s.content = content

	return s
}

func (s *TestService) ReadContent() string {
	return s.content
}
