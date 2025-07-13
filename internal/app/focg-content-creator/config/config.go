// Package config provides configuration structures and constants for the FOCG Content Creator application.
package config

import "path"

// Config defines the structure for the application's configuration file.
type Config struct {
	Version   string `json:"version"`              // Version of the config, e.g., "1.0.0"
	Port      string `json:"port"`                 // Port for the web server, default is ":3621"
	Author    string `json:"author,omitempty"`     // Optional field for author name
	AuthorURL string `json:"author_url,omitempty"` // Optional field for author URL
}

// configPath is the relative path to the application's configuration file.
var configPath = path.Join("data", "config.json")
