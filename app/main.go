package main

import (
	"html/template"
	"io"

	"github.com/kevcal69/go-redis-email/pkg/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func registerAPIHandlers(e *echo.Echo) {
	for _, view := range handlers.APIRoutes {
		e.Match(view.Method, "api"+view.Route, view.Handler)
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRFToken",
		CookieName:  "csrftoken",
	}))
	e.Static("/static", "pkg/templates/public")

	// Define API routes
	registerAPIHandlers(e)
	// Define the vuejs routes
	e.File("/*", "pkg/templates/public/index.html")

	// Start the server
	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
