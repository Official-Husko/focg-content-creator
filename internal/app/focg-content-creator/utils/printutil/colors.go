// Package printutil provides ANSI color constants for terminal output formatting.
package printutil

// ANSI color escape codes for use in terminal output.
const (
	ColorGray   = "\033[90m" // Gray (bright black)
	ColorPurple = "\033[35m" // Purple (magenta)
	ColorCyan   = "\033[36m" // Cyan
	ColorReset  = "\033[0m"  // Reset to default color
	ColorWarn   = "\033[33m" // Yellow (warning)
	ColorError  = "\033[31m" // Red (error)
	ColorInfo   = "\033[32m" // Green (info/success)
)
