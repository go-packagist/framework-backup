package gin

import "github.com/go-packagist/framework/foundation"

type Provider struct {
	app *foundation.Application
}

var _ foundation.Provider = (*Provider)(nil)

func NewGinProvider(app *foundation.Application) foundation.Provider {
	return &Provider{
		app: app,
	}
}

func (p Provider) Register() {
	p.app.Singleton("gin", func(app *foundation.Application) interface{} {
		return NewGin(app)
	})
}
