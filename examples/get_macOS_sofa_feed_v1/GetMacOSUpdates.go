// GetMacOSUpdates.go
// Fetches and displays full macOS update data using the SOFA v1 feed.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-sofa/client"
	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/sofa"
)

func main() {
	fmt.Println("Fetching macOS updates from v1 API...")

	httpClient, err := client.NewV1Client()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer httpClient.Close()

	c := sofa.NewClient(httpClient)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := c.GetMacOSSofaFeedV1(ctx)
	if err != nil {
		log.Fatalf("Error fetching macOS updates: %v", err)
	}

	jsonOutput, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON output: %v", err)
	}

	fmt.Println(string(jsonOutput))
}
