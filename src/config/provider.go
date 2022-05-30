package config

import "github.com/go-packagist/foundation"

type Provider struct {
	app *foundation.Application
}

var _ foundation.Provider = (*Provider)(nil)

func NewConfigProvider(app *foundation.Application) foundation.Provider {
	return &Provider{
		app: app,
	}
}

func (p *Provider) Register() {
	p.app.Instance("config", New(&Options{}))
}
