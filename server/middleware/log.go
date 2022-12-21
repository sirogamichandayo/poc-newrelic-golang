package middleware

import (
	"fmt"
	"github.com/dijsilva/golang-api-newrelic/pkg"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func RegisterLogger(app *newrelic.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		if app != nil {

			fmt.Println(pkg.NewLoggerWithContext(c, app))
			c.Set(
				pkg.LoggerContextKey,
				pkg.NewLoggerWithContext(c, app),
			)
		}
		c.Next()
	}
}
