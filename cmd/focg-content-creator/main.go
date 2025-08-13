package main

import (
	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/webserver"
)

func main() {
	// Start the web server in a goroutine
	go webserver.Start()

	select {}
}
