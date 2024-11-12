package main

import (
	"net/http"

	"github.com/aiosifelis/go-web-app-template/src/middlewares"
	"github.com/aiosifelis/go-web-app-template/src/migrations"
	"github.com/aiosifelis/go-web-app-template/src/system"
	"github.com/aiosifelis/go-web-app-template/src/views/layouts"
	"github.com/aiosifelis/go-web-app-template/src/views/pages"
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

	db, err := system.InitializeDatabase()
	if err != nil {
		panic("connection")
	}

	// Migrate the database
	migrations.Migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.Database(db))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.Logger().Error(err)
		// c.Render(http.StatusInternalServerError, "error.html", map[string]interface{}{
		// 	"Title":   "Error Page",
		// 	"Message": "An error occurred",
		// })
		c.Redirect(http.StatusFound, "/error")
	}

	e.Static("public", "src/public")

	authRoutes := e.Group("/auth")
	mainRoutes := e.Group("/")

	mainRoutes.Use(middlewares.AuthMiddleware)

	mainRoutes.GET("/about", func(c echo.Context) error {
		return system.RenderTemplate(c, pages.AboutPage(&layouts.PageData{
			Title:   "About Page",
			Content: "This is the about page",
		}))
	})

	authRoutes.GET("/login", func(c echo.Context) error {
		return system.RenderTemplate(c, pages.LoginPage(&layouts.PageData{
			Title:   "Login Page",
			Content: "This is the login page",
		}))
	})

	authRoutes.GET("/signup", func(c echo.Context) error {
		return system.RenderTemplate(c, pages.SignupPage(&layouts.PageData{
			Title:   "Signup Page",
			Content: "This is the signup page",
		}))
	})

	authRoutes.GET("/forgot-password", func(c echo.Context) error {
		return system.RenderTemplate(c, pages.ForgotPasswordPage(&layouts.PageData{
			Title:   "Forgot Password Page",
			Content: "This is the forgot password page",
		}))
	})

	e.GET("/", func(c echo.Context) error {
		return system.RenderTemplate(c, pages.HomePage(&layouts.PageData{
			Title:   "Home Page",
			Content: "This is the home page",
		}))
	})

	e.GET("/error", func(c echo.Context) error {
		return system.RenderTemplate(c, pages.ErrorPage(&layouts.PageData{
			Title:   "Error Page",
			Content: "This is the error page",
		}))
	})

	e.Logger.Fatal(e.Start(":3000"))
}
