package routes

import (
	"html/template"
	"net/http"
	"path"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
	"github.com/go-chi/chi/v5"
)

func RegisterTraitRoute(r chi.Router, templatesDir string) {
	r.Get("/trait", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "trait.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse trait template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Trait Editor - FoC:G Content Creator",
			"Header":     "Trait Editor",
			"Message":    "This is the Trait Editor.",
			"ActivePage": "trait",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute trait template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
