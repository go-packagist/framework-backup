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
	return
}
