package webserver

import (
	"html/template"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
)

func RegisterRoutes(r chi.Router, templatesDir string) {
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

	r.Get("/generator", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "generator.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse generator template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Generator - FoC:G Content Creator",
			"Header":     "Generator",
			"Message":    "This is the generator page.",
			"ActivePage": "generator",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute generator template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

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

	r.Get("/race", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "race.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse race template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Race Editor - FoC:G Content Creator",
			"Header":     "Race Editor",
			"Message":    "This is the Race Editor.",
			"ActivePage": "race",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute race template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

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

	r.Get("/translation", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			path.Join(templatesDir, "index.html"),
			path.Join(templatesDir, "components/sidebar.html"),
			path.Join(templatesDir, "translation.html"),
		)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to parse translation template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Title":      "Translation Editor - FoC:G Content Creator",
			"Header":     "Translation Editor",
			"Message":    "This is the Translation Editor.",
			"ActivePage": "translation",
		}
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			printutil.Print("ERROR", "WebServer", "Failed to execute translation template: "+err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
