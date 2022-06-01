package gin

import (
	g "github.com/gin-gonic/gin"
	"github.com/go-packagist/framework/foundation"
)

func NewGin(app *foundation.Application) *g.Engine {
	return g.Default()
}
