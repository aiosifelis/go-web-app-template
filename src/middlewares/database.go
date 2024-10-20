package middlewares

import (
	"github.com/aiosifelis/go-web-app-template/src/system"
	"github.com/labstack/echo/v4"
)

// get the pool passed to the middleware
func Database(db *system.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
