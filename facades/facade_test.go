package facades

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFacade_App(t *testing.T) {
	foundation.NewApplication("./")

	// assert.Equal(t, "./app", App().getPath())
	assert.Equal(t, "./app", App().Container.MustMake("path").(string))
}
