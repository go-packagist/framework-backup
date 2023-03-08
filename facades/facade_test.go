package facades

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/hashing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createApp() *foundation.Application {
	return foundation.NewApplication()
}

func TestFacades_Hashing(t *testing.T) {
	app := createApp()

	app.Register(hashing.NewHashProvider(app.Container))

	value := "123456"

	hashedValue1 := MustHash().Driver().MustMake(value)
	assert.True(t, MustHash().Driver().Check(value, hashedValue1))
}
