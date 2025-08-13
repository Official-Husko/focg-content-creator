package routes

import (
	"html/template"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
)

func RegisterSubraceRoute(r chi.Router, templatesDir string) {
	r.Get("/subrace", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "subrace.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse subrace template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Subrace Editor - FoC:G Content Creator",
			"Header":     "Subrace Editor",
			"Message":    "This is the Subrace Editor.",
			"ActivePage": "subrace",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute subrace template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
