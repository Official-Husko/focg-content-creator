// Package printutil provides utilities for colorized and formatted terminal output.
package printutil

import "time"

// FormatTimestamp returns the current local time as a string in the format DD/MM/YYYY HH:MM:SS.
// This format is used for log message timestamps.
func FormatTimestamp() string {
	return time.Now().Format("02/01/2006 15:04:05")
}
