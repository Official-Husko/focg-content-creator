// Package config provides functions for loading application configuration from disk.
package config

import (
	"encoding/json"
	"log"
	"os"
)

// LoadOrCreateConfig loads the application's configuration from disk, or creates a default config if missing.
//
// If the config file does not exist, a default config is written and returned.
// If the config file exists but cannot be opened or decoded, the program will log a fatal error and exit.
//
// Returns:
//   - Config struct loaded from disk, or the default config if the file was missing.
func LoadOrCreateConfig() Config {
	defaultConfig := Config{Port: ":8080"}
	file, err := os.Open(configPath)
	if os.IsNotExist(err) {
		data, _ := json.MarshalIndent(defaultConfig, "", "  ")
		_ = os.WriteFile(configPath, data, 0644)
		return defaultConfig
	} else if err != nil {
		log.Fatalf("Error opening config: %v", err)
	}
	defer file.Close()
	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatalf("Error decoding config: %v", err)
	}
	return cfg
}
