package routes

import (
	"html/template"
	"net/http"
	"path"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
	"github.com/go-chi/chi/v5"
)

func RegisterTextsRoute(r chi.Router, templatesDir string) {
	r.Get("/texts", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "texts.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse texts template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Texts Editor - FoC:G Content Creator",
			"Header":     "Texts Editor",
			"Message":    "This is the Texts Editor.",
			"ActivePage": "texts",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute texts template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
