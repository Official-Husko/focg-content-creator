package routes

import (
	"html/template"
	"net/http"
	"path"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
	"github.com/go-chi/chi/v5"
)

func RegisterQuestRoute(r chi.Router, templatesDir string) {
	r.Get("/quest", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "quest.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse quest template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Quest Editor - FoC:G Content Creator",
			"Header":     "Quest Editor",
			"Message":    "This is the Quest Editor.",
			"ActivePage": "quest",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute quest template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
