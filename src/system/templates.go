package system

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return echo.ErrNotFound
	}

	return tmpl.ExecuteTemplate(w, "index.html", data)
}

func NewTemplateRenderer() *Template {

	viewsDir := "src/views"

	tmpl := make(map[string]*template.Template)

	// Walk through the views directory and parse templates dynamically
	filepath.Walk(viewsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Only parse .html files
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			name := strings.TrimPrefix(path, viewsDir+"/")
			// Parse the file along with base.html
			tmpl[name] = template.Must(template.ParseFiles(path, filepath.Join(viewsDir, "index.html")))
		}
		return nil
	})

	fmt.Println("Templates Loaded")

	return &Template{
		templates: tmpl,
	}
}
