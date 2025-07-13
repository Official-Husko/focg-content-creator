// Package focgcontentcreator provides name generation logic for various fantasy cultures.
package focgcontentcreator

import (
	"math/rand"
)

// GenerateNames produces a list of generated names using the provided NameGenData and rules for a given gender or type.
//
// Parameters:
//   - data: Pointer to the loaded NameGenData struct containing parts and rules.
//   - gender: The key for the rules map (e.g., "male", "female").
//   - count: The number of names to generate.
//
// Returns:
//   - A slice of generated names as strings. If no rules exist for the gender, returns nil.
func GenerateNames(data *NameGenData, gender string, count int) []string {
	rules, ok := data.Rules[gender]
	if !ok || len(rules) == 0 {
		return nil // No rules for the specified gender/type
	}
	// Preallocate the result slice for efficiency
	names := make([]string, 0, count)
	for i := 0; i < count; i++ {
		// Select a random rule for this gender/type
		rule := rules[rand.Intn(len(rules))]
		var name string
		for _, step := range rule {
			if arr, ok := data.Parts[step]; ok && len(arr) > 0 {
				// If the step is a part, pick a random value from it
				name += arr[rand.Intn(len(arr))]
			} else {
				// Otherwise, treat the step as a literal (e.g., space, dash, etc.)
				name += step
			}
		}
		names = append(names, name)
	}
	return names
}
