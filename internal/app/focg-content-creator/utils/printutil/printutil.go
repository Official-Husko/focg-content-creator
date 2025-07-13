// Package printutil provides formatted and colorized printing utilities for terminal output.
package printutil

import (
	"fmt"
)

// Print outputs a formatted, colorized log message with timestamp, type, module, and message.
//
// msgType: Log type (e.g., "INFO", "WARN", "ERROR").
// module:  Name of the module emitting the message.
// message: The message to print.
// colorMap: (optional) map of substrings in message to color constants for highlighting.
//
// Example:
//
//	printutil.Print("INFO", "WebServer", "Started at http://localhost:8080", map[string]string{"http://localhost:8080": printutil.ColorInfo})
func Print(msgType, module, message string, colorMap ...map[string]string) {
	timestamp := FormatTimestamp()
	msg := message
	if len(colorMap) > 0 {
		msg = Colorize(message, colorMap[0])
	}
	// Print: [ and ] in gray, module in purple, message in cyan, type colored, timestamp gray
	fmt.Printf("%s%s%s %s%-5s%s%s[%s%s%s]%s %s%s%s\n",
		ColorGray, timestamp, ColorReset,
		colorForType(msgType), msgType, ColorReset,
		ColorGray, ColorPurple, module, ColorGray, ColorReset, // [module] with [] gray, text purple
		ColorCyan, msg, ColorReset,
	)
}

// colorForType returns the color constant for a given log type.
func colorForType(msgType string) string {
	switch msgType {
	case "WARN":
		return ColorWarn
	case "ERROR":
		return ColorError
	case "INFO":
		return ColorInfo
	default:
		return ColorGray
	}
}
