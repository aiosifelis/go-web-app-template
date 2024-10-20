package middlewares

import (
	"github.com/aiosifelis/go-web-app-template/src/system"
	"github.com/labstack/echo/v4"
)

func Database(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := system.InitDB("root:password@tcp(localhost:3306)/go_web_app")
		if err != nil {
			return err
		}
		defer db.Close()
		c.Set("db", db)
		return next(c)
	}
}
