package main

import (
	"github.com/Official-Husko/focg-content-creator/internal/app/focg-content-creator/webserver"
)

func main() {
	// Start the web server in a goroutine
	go webserver.Start()

	/*
		jsonPath := filepath.FromSlash(name_generator.NordNameGenPath)
		data, err := name_generator.LoadNameGenData(jsonPath)
		if err != nil {
			log.Fatalf("Failed to load name data: %v", err)
		}
		fmt.Println("Sample fantasy nord male names:")
		for _, name := range name_generator.GenerateNames(data, "male", 10) {
			fmt.Println(name)
		}
		fmt.Println("\nSample fantasy nord female names:")
		for _, name := range name_generator.GenerateNames(data, "female", 10) {
			fmt.Println(name)
		}
		// Keep the main goroutine alive
	*/
	select {}
}
