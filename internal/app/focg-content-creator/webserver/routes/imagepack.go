package routes

import (
	"html/template"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
)

func RegisterImagepackRoute(r chi.Router, templatesDir string) {
	r.Get("/imagepack", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "imagepack.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse imagepack template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Imagepack Editor - FoC:G Content Creator",
			"Header":     "Imagepack Editor",
			"Message":    "This is the Imagepack Editor.",
			"ActivePage": "imagepack",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute imagepack template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
