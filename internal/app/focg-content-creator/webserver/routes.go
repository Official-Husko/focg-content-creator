package webserver

import (
	"github.com/go-chi/chi/v5"

	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/webserver/routes"
)

func RegisterRoutes(r chi.Router, templatesDir string) {
	routes.RegisterIndexRoute(r, templatesDir)
	routes.RegisterNameGenRoute(r, templatesDir)
	routes.RegisterSettingsRoute(r, templatesDir)
	routes.RegisterImagepackRoute(r, templatesDir)
	routes.RegisterRaceRoute(r, templatesDir)
	routes.RegisterSubraceRoute(r, templatesDir)
	routes.RegisterTraitRoute(r, templatesDir)
	routes.RegisterQuestRoute(r, templatesDir)
	routes.RegisterTextsRoute(r, templatesDir)
	routes.RegisterTranslationRoute(r, templatesDir)
	routes.RegisterSubraceRoute(r, templatesDir)
}
