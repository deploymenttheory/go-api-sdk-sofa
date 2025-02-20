// macos_updates_client.go
// SOFA | A MacAdmin's Simple Organized Feed for Apple Software Updates
// API reference: https://sofafeed.macadmins.io/v1/macos_data_feed.json
// This client retrieves macOS update data in JSON format.

package sofamacos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// API Endpoint
const uriAPIMacOSUpdates = "https://sofafeed.macadmins.io/v1/macos_data_feed.json"

// Root represents the top-level structure of the macOS update feed response
type Root struct {
	UpdateHash          string                  `json:"UpdateHash"`
	OSVersions          []OSVersion             `json:"OSVersions"`
	XProtectPayloads    XProtectPayloads        `json:"XProtectPayloads"`
	XProtectPlistConfig XProtectPlistConfigData `json:"XProtectPlistConfigData"`
	Models              map[string]Model        `json:"Models"`
	InstallationApps    InstallationApps        `json:"InstallationApps"`
}

// OSVersion represents an operating system version
type OSVersion struct {
	OSVersion        string            `json:"OSVersion"`
	Latest           ReleaseInfo       `json:"Latest"`
	SecurityReleases []SecurityRelease `json:"SecurityReleases"`
	SupportedModels  []SupportedModel  `json:"SupportedModels"`
}

// ReleaseInfo represents the latest release information
type ReleaseInfo struct {
	ProductVersion        string          `json:"ProductVersion"`
	Build                 string          `json:"Build"`
	ReleaseDate           string          `json:"ReleaseDate"`
	ExpirationDate        string          `json:"ExpirationDate"`
	SupportedDevices      []string        `json:"SupportedDevices"`
	SecurityInfo          string          `json:"SecurityInfo"`
	CVEs                  map[string]bool `json:"CVEs"`
	ActivelyExploitedCVEs []string        `json:"ActivelyExploitedCVEs"`
	UniqueCVEsCount       int             `json:"UniqueCVEsCount"`
}

// SecurityRelease represents a security update
type SecurityRelease struct {
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

// SupportedModel represents Mac models supported for an OS version
type SupportedModel struct {
	Model       string            `json:"Model"`
	URL         string            `json:"URL"`
	Identifiers map[string]string `json:"Identifiers"`
}

// XProtectPayloads represents security update configurations
type XProtectPayloads struct {
	XProtectFramework string `json:"com.apple.XProtectFramework.XProtect"`
	PluginService     string `json:"com.apple.XprotectFramework.PluginService"`
	ReleaseDate       string `json:"ReleaseDate"`
}

// XProtectPlistConfigData represents the XProtect configuration data
type XProtectPlistConfigData struct {
	XProtect    string `json:"com.apple.XProtect"`
	ReleaseDate string `json:"ReleaseDate"`
}

// Model represents a Mac model with supported OS versions
type Model struct {
	MarketingName string   `json:"MarketingName"`
	SupportedOS   []string `json:"SupportedOS"`
	OSVersions    []int    `json:"OSVersions"`
}

// InstallationApps represents available macOS installation apps
type InstallationApps struct {
	LatestUMA      AppEntry   `json:"LatestUMA"`
	AllPreviousUMA []AppEntry `json:"AllPreviousUMA"`
	LatestMacIPSW  MacIPSW    `json:"LatestMacIPSW"`
}

// AppEntry represents an application entry for macOS installation
type AppEntry struct {
	Title     string `json:"title"`
	Version   string `json:"version"`
	Build     string `json:"build"`
	AppleSlug string `json:"apple_slug"`
	URL       string `json:"url"`
}

// MacIPSW represents the latest macOS IPSW details
type MacIPSW struct {
	MacOSIPSWURL       string `json:"macos_ipsw_url"`
	MacOSIPSWBuild     string `json:"macos_ipsw_build"`
	MacOSIPSWVersion   string `json:"macos_ipsw_version"`
	MacOSIPSWAppleSlug string `json:"macos_ipsw_apple_slug"`
}

// GetMacOSUpdates fetches macOS update data from the API
func GetMacOSUpdates() (*Root, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(uriAPIMacOSUpdates)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS updates: %w", err)
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
