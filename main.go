package main

import (
	"net/http"

	"github.com/aiosifelis/go-web-app-template/src/system"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("public", "src/public")
	e.Renderer = system.TemplateRenderer()

	e.GET("/", func(c echo.Context) error {

		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"name": "World-eee",
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
