package foundation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplication(t *testing.T) {
}

func TestApplication_Version(t *testing.T) {
	app := NewApplication("./")

	assert.Equal(t, "0.0.1", app.Version())
}

func TestApplication_Register(t *testing.T) {
	app := NewApplication("./")

	app.Register("test", NewTestProvider(app))

	assert.Equal(t, map[string]Provider{
		"test": NewTestProvider(app),
	}, app.GetProviders())
}
