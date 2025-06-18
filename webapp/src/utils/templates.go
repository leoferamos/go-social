package utils

import (
	"io"
	"text/template"
)

var templates *template.Template

// LoadTemplates loads all HTML templates from the views directory.
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/partials/*.html"))
}

// ExecuteTemplate executes the specified template with the provided data.
func ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	if templates == nil {
		LoadTemplates()
	}

	return templates.ExecuteTemplate(w, name, data)
}
