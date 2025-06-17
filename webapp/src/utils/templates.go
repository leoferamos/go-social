package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadTemplates loads all HTML templates from the views directory.
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/partials/*.html"))
}

// ExecuteTemplate executes the specified template with the provided data.
func ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) error {
	if templates == nil {
		LoadTemplates()
	}

	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}
	return nil
}
