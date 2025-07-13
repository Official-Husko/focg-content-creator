// Package config provides functions for writing application configuration to disk.
package config

import (
	"encoding/json"
	"os"
	"path"
)

// WriteConfig writes the provided Config struct to disk as JSON.
//
// If the config file or its parent directory does not exist, they will be created.
// Returns an error if writing fails.
func WriteConfig(cfg Config) error {
	if err := os.MkdirAll(path.Dir(configPath), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

// TODO: Implement a queue system to handle config writes in the background to avoid blocking the main thread
// TODO: Write using sjson and gjson to allow partial updates without rewriting the entire file
// TODO: Create config migrations to handle changes in config structure over time
