// Package webserver provides the HTTP server for serving the FOCG Content Creator web UI.
package webserver

import (
	"net/http"
	"os"
	"path"

	"github.com/go-chi/chi/v5"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/config"
	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/utils/printutil"
)

// getWebDir returns the path to the web directory, using the WEB_DIR environment variable if set.
func getWebDir() string {
	if dir := os.Getenv("WEB_DIR"); dir != "" {
		return dir
	}
	return path.Join("..", "..", "web")
}

var webDir = getWebDir()
var templatesDir = path.Join(webDir, "templates")

// Start launches the web server to serve the web UI using the chi router.
//
// Loads configuration, checks for the web directory, and starts the HTTP server.
// Prints status and error messages using the printutil package.
func Start() {
	cfg := config.LoadOrCreateConfig()
	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		printutil.Print("ERROR", "WebServer", "Web directory '"+webDir+"' does not exist")
		os.Exit(1)
	}

	r := chi.NewRouter()

	// Register all web routes in a separate function
	RegisterRoutes(r, templatesDir)

	// Serve static files only under /static/
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(webDir))))

	printutil.Print(
		"INFO", "WebServer",
		"Serving FoC:G Content Creator at http://localhost"+cfg.Port+" ...",
		map[string]string{"http://localhost" + cfg.Port: printutil.ColorInfo},
	)
	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		printutil.Print("ERROR", "WebServer", err.Error())
		os.Exit(1)
	}
}
