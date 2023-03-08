package hashing

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	app := createApp()

	app.Register(NewHashProvider(app.Container))

	manager := app.MustMake("hash").(*Manager)
	value := "123456"

	// hash(default:md5)
	hashedValue := manager.Driver().MustMake(value)
	assert.True(t, manager.Driver().Check(value, hashedValue))

	// hash bcrypt
	hashedValue2 := manager.Driver("bcrypt").MustMake(value)
	assert.True(t, manager.Driver("bcrypt").Check(value, hashedValue2))

	// hash md5
	hashedValue3 := manager.Driver("md5").MustMake(value)
	assert.True(t, manager.Driver("md5").Check(value, hashedValue3))
}

func createApp() *foundation.Application {
	return foundation.NewApplication()
}
