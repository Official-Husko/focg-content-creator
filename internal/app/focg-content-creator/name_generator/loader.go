// Package focgcontentcreator provides loading utilities for name generator data files.
package focgcontentcreator

import (
	"encoding/json"
	"os"
)

// LoadNameGenData loads and parses a NameGenData struct from a JSON file at the given path.
//
// Parameters:
//   - path: The file path to the JSON data file.
//
// Returns:
//   - Pointer to a NameGenData struct if successful, or an error if loading or parsing fails.
func LoadNameGenData(path string) (*NameGenData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}
	var data NameGenData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
