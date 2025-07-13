// Package focgcontentcreator provides data structures and constants for the name generator system.
package focgcontentcreator

import "path/filepath"

// NameGenData defines the structure for a generic name generator JSON file.
// This structure is compatible with various name generator types (e.g., nord, elven).
type NameGenData struct {
	Version string `json:"version"`
	// Name is the display name of the name generator (e.g., "Nord Names").
	Name string `json:"name"`
	// Description provides a human-readable description of the generator.
	Description string `json:"description"`
	// Parts maps part names (e.g., "prefix", "suffix") to lists of possible values.
	Parts map[string][]string `json:"parts"`
	// Rules maps gender or type (e.g., "male", "female") to a list of rules, where each rule is a sequence of part names.
	Rules map[string][][]string `json:"rules"`
	// Source is an optional field indicating the data's origin or reference.
	Source string `json:"source"`
}

// NordNameGenPath is the relative path to the default Nord name generator JSON data file.
var NordNameGenPath = filepath.Join("..", "..", "internal", "data", "name_generator", "nord.json")
