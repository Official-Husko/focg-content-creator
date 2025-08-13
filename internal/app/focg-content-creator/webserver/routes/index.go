package routes

import (
	"html/template"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
)

func RegisterIndexRoute(r chi.Router, templatesDir string) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "FoC:G Content Creator",
			"Header":     "Welcome to FoC:G Content Creator",
			"Message":    "This is the web interface for your mod/content creation tool.",
			"ActivePage": "home",
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
