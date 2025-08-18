package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, template_name string, data any) {
	base_files := []string{
		"templates/base.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/alerts.html",
	}

	base_files = append(base_files, fmt.Sprintf("templates/%s", template_name))

	t := template.Must(template.ParseFiles(base_files...))

	t.ExecuteTemplate(w, "base", data)
}
