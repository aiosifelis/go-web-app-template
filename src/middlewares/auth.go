package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// create a random number
		// if the number is even, return 401
		// if the number is odd, continue to the next middleware
		// this is a dummy example of an authentication middleware

		if 1 == 1 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}

		return next(c)
	}
}
