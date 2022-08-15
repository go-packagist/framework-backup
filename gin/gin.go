package gin

import (
	g "github.com/gin-gonic/gin"
	"github.com/go-packagist/framework/container"
)

func NewGin(c *container.Container) *g.Engine {
	return g.Default()
}
