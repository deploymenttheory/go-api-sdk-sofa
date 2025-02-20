// main_ios.go
// Fetches iOS update data using the SOFA feed and exports to JSON file.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/sofamacos"
)

func main() {
	fmt.Println("Fetching iOS updates...")

	data, err := sofamacos.GetMacOSUpdates()
	if err != nil {
		log.Fatalf("Error fetching iOS updates: %v", err)
	}

	// Create filename with timestamp
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("ios_updates_%s.json", timestamp)

	// Marshal JSON with indentation
	jsonOutput, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON output: %v", err)
	}

	// Write to file
	err = os.WriteFile(filename, jsonOutput, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Printf("Successfully exported iOS updates to %s\n", filename)
}
