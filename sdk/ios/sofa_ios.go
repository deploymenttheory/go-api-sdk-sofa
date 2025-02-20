// ios_updates_client.go
// SOFA | A MacAdmin's Simple Organized Feed for Apple Software Updates
// API reference: https://sofafeed.macadmins.io/v1/ios_data_feed.json
// This client retrieves iOS update data in JSON format.

package sofaios

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// API Endpoint
const uriAPIiOSUpdates = "https://sofafeed.macadmins.io/v1/ios_data_feed.json"

// Root represents the top-level structure of the iOS update feed response
type Root struct {
	UpdateHash string      `json:"UpdateHash"`
	OSVersions []OSVersion `json:"OSVersions"`
}

// OSVersion represents an iOS version entry
type OSVersion struct {
	OSVersion        string           `json:"OSVersion"`
	Latest           ReleaseInfo      `json:"Latest"`
	SecurityReleases []SecurityUpdate `json:"SecurityReleases"`
}

// ReleaseInfo represents the latest release details
type ReleaseInfo struct {
	ProductVersion        string          `json:"ProductVersion"`
	Build                 string          `json:"Build"`
	ReleaseDate           string          `json:"ReleaseDate"`
	ExpirationDate        string          `json:"ExpirationDate"`
	SupportedDevices      []string        `json:"SupportedDevices"`
	SecurityInfo          string          `json:"SecurityInfo"`
	UniqueCVEsCount       int             `json:"UniqueCVEsCount"`
	ActivelyExploitedCVEs []string        `json:"ActivelyExploitedCVEs"`
	CVEs                  map[string]bool `json:"CVEs"`
}

// SecurityUpdate represents security updates for an iOS version
type SecurityUpdate struct {
	UpdateName               string          `json:"UpdateName"`
	ProductName              string          `json:"ProductName"`
	ProductVersion           string          `json:"ProductVersion"`
	ReleaseDate              string          `json:"ReleaseDate"`
	ReleaseType              string          `json:"ReleaseType"`
	SecurityInfo             string          `json:"SecurityInfo"`
	SupportedDevices         []string        `json:"SupportedDevices"`
	CVEs                     map[string]bool `json:"CVEs"`
	ActivelyExploitedCVEs    []string        `json:"ActivelyExploitedCVEs"`
	UniqueCVEsCount          int             `json:"UniqueCVEsCount"`
	DaysSincePreviousRelease int             `json:"DaysSincePreviousRelease"`
}

// GetiOSUpdates fetches iOS update data from the API
func GetiOSUpdates() (*Root, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(uriAPIiOSUpdates)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch iOS updates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var data Root
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &data, nil
}
