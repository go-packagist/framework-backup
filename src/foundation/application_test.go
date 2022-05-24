package foundation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplication(t *testing.T) {
}

func TestApplication_Version(t *testing.T) {
	app := NewApplication("./")

	assert.Equal(t, VERSION, app.Version())
}

func TestApplication_Register(t *testing.T) {
	app := NewApplication("./")

	app.Register("test", NewTestProvider(app))

	assert.Equal(t, map[string]Provider{
		"test": NewTestProvider(app),
	}, app.GetProviders())

	// Test Service
	testService := app.GetService("test").(*TestService)
	assert.Equal(t, app, testService.Application())

	testService.WriteContent("123123")
	assert.Equal(t, "123123", testService.ReadContent())

	testService.WriteContent("234567")
	assert.Equal(t, "234567", testService.ReadContent())

	app.GetService("test").(*TestService).WriteContent("123123")
	fmt.Println(app.GetService("test").(*TestService).ReadContent())
	fmt.Println(testService.ReadContent())

	fmt.Println(app)
}
