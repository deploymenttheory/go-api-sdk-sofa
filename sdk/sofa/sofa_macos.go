// sofa_macos.go
// SOFA | A MacAdmin's Simple Organized Feed for Apple Software Updates
// API reference: https://sofafeed.macadmins.io/v1/macos_data_feed.json
// This client retrieves macOS update data in JSON format.

package sofa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/models"
)

// API Endpoint
const uriAPIMacOSUpdates = "https://sofafeed.macadmins.io/v1/macos_data_feed.json"

// ResourceSofamacOSRoot represents the top-level structure of the macOS update feed response
type ResourceSofamacOSRoot struct {
	UpdateHash          string                       `json:"UpdateHash"`
	OSVersions          []macOSVersion               `json:"OSVersions"`
	XProtectPayloads    macOSXProtectPayloads        `json:"XProtectPayloads"`
	XProtectPlistConfig macOSXProtectPlistConfigData `json:"XProtectPlistConfigData"`
	Models              map[string]macOSModel        `json:"Models"`
	InstallationApps    InstallationApps             `json:"InstallationApps"`
	SupportedModels     []macOSSupportedModel        `json:"SupportedModels"`
}

// macOSVersion  represents an operating system version
type macOSVersion struct {
	OSVersion        string                 `json:"OSVersion"`
	Latest           macOSReleaseInfo       `json:"Latest"`
	SecurityReleases []macOSSecurityRelease `json:"SecurityReleases"`
	SupportedModels  []macOSSupportedModel  `json:"SupportedModels"`
}

// macOSReleaseInfo represents the latest release information
type macOSReleaseInfo struct {
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

// macOSSecurityRelease represents a security update
type macOSSecurityRelease struct {
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

// macOSSupportedModel represents Mac models supported for an OS version
type macOSSupportedModel struct {
	Model       string            `json:"Model"`
	URL         string            `json:"URL"`
	Identifiers map[string]string `json:"Identifiers"`
}

// macOSXProtectPayloads represents security update configurations
type macOSXProtectPayloads struct {
	XProtectFramework string `json:"com.apple.XProtectFramework.XProtect"`
	PluginService     string `json:"com.apple.XprotectFramework.PluginService"`
	ReleaseDate       string `json:"ReleaseDate"`
}

// XProtectPlistConfigData represents the XProtect configuration data
type macOSXProtectPlistConfigData struct {
	XProtect    string `json:"com.apple.XProtect"`
	ReleaseDate string `json:"ReleaseDate"`
}

// macOSModel represents a Mac model with supported OS versions
type macOSModel struct {
	MarketingName string   `json:"MarketingName"`
	SupportedOS   []string `json:"SupportedOS"`
	OSVersions    []int    `json:"OSVersions"`
}

// macOSInstallationApps represents available macOS installation apps
type InstallationApps struct {
	LatestUMA      AppEntry   `json:"LatestUMA"`
	AllPreviousUMA []AppEntry `json:"AllPreviousUMA"`
	LatestMacIPSW  MacIPSW    `json:"LatestMacIPSW"`
}

// macOSAppEntry represents an application entry for macOS installation
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
func GetMacOSUpdates() (*models.MacOSFeedResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(uriAPIMacOSUpdates)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS updates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var data models.MacOSFeedResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &data, nil
}
