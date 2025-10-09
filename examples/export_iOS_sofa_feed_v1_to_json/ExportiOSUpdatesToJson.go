// ExportiOSUpdatesToJson.go
// Fetches iOS update data using the SOFA v1 feed and exports to JSON file.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-sdk-sofa/client"
	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/sofa"
)

func main() {
	fmt.Println("Fetching iOS updates from v1 API...")

	httpClient, err := client.NewV1Client()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer httpClient.Close()

	c := sofa.NewClient(httpClient)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := c.GetiOSSofaFeedV1(ctx)
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
