package routes

import (
	"html/template"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
)

func RegisterSettingsRoute(r chi.Router, templatesDir string) {
	r.Get("/settings", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "settings.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse settings template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Settings - FoC:G Content Creator",
			"Header":     "Settings",
			"Message":    "This is the settings page.",
			"ActivePage": "settings",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute settings template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
