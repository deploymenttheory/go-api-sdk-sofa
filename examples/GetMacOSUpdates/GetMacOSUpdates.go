// main_ios.go
// Fetches and displays full iOS update data using the SOFA feed.

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/sofamacos"
)

func main() {
	fmt.Println("Fetching iOS updates...")

	data, err := sofamacos.GetMacOSUpdates()
	if err != nil {
		log.Fatalf("Error fetching iOS updates: %v", err)
	}

	jsonOutput, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON output: %v", err)
	}

	fmt.Println(string(jsonOutput))
}
