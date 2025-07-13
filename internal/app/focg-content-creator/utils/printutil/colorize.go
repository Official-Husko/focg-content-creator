// Package printutil provides utilities for colorized and formatted terminal output.
package printutil

import (
	"regexp"
)

// Colorize returns a copy of text with specified substrings wrapped in ANSI color codes.
//
// The colorMap argument maps substrings to color constants (e.g., ColorCyan).
// Longer substrings are replaced first to avoid partial overlaps.
//
// Example:
//
//	Colorize("Hello world", map[string]string{"world": ColorCyan})
//
// Returns:
//
//	The colorized string, with all specified substrings wrapped in their color and ColorReset.
func Colorize(text string, colorMap map[string]string) string {
	if len(colorMap) == 0 {
		return text
	}
	// Sort keys by length descending to avoid partial overlaps
	keys := make([]string, 0, len(colorMap))
	for k := range colorMap {
		keys = append(keys, k)
	}
	// Sort by length descending
	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			if len(keys[j]) > len(keys[i]) {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}
	result := text
	for _, k := range keys {
		c := colorMap[k]
		// Use regexp to replace all occurrences, escaping special chars
		re := regexp.MustCompile(regexp.QuoteMeta(k))
		result = re.ReplaceAllString(result, c+k+ColorReset)
	}
	return result
}
