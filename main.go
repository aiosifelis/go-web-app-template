package main

import (
	"net/http"

	"github.com/aiosifelis/go-web-app-template/src/middlewares"
	"github.com/aiosifelis/go-web-app-template/src/system"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	connectionPool, err := system.CreateConnectionPool()
	if err != nil {
		panic("connection")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = system.NewTemplateRenderer()
	e.Use(middlewares.Database(&connectionPool))

	e.Static("public", "src/public")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{
			"Title":   "Home Page",
			"Message": "Welcome to the Home Page",
		})
	})

	e.GET("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about.html", map[string]interface{}{
			"Title":   "About Page",
			"Message": "Learn more About Us on this page.",
		})
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "auth/login.html", map[string]interface{}{
			"Title":   "Login Page",
			"Message": "Login to your account",
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
