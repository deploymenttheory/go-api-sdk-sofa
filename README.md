# Go API SDK for SOFA

A Go SDK for interacting with the SOFA (Simple Organized Feed for Apple) API, providing access to macOS and iOS update information.

[![Go Version](https://img.shields.io/github/go-mod/go-version/deploymenttheory/go-api-sdk-sofa)](https://go.dev/)
[![License](https://img.shields.io/github/license/deploymenttheory/go-api-sdk-sofa)](LICENSE)

## About SOFA

SOFA (Simple Organized Feed for Apple) supports MacAdmins by efficiently tracking and surfacing information on updates for macOS and iOS. It consists of a machine-readable feed and user-friendly web interface, providing continuously up-to-date information on XProtect data, OS updates, and the details bundled in those releases.

## About This SDK

This Go SDK provides a simple, idiomatic way to access the SOFA API feeds for macOS and iOS updates. It handles the HTTP requests, JSON parsing, and provides strongly-typed structures for working with the data.

## Installation

```bash
go get github.com/deploymenttheory/go-api-sdk-sofa
```

## Usage

### Fetching macOS Updates

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/sofa"
)

func main() {
	fmt.Println("Fetching macOS updates...")

	data, err := sofa.GetMacOSUpdates()
	if err != nil {
		log.Fatalf("Error fetching macOS updates: %v", err)
	}

	// Access structured data
	fmt.Printf("Latest macOS: %s (Build %s)\n", 
		data.OSVersions[0].Latest.ProductVersion, 
		data.OSVersions[0].Latest.Build)

	// Or convert to JSON for display/storage
	jsonOutput, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON output: %v", err)
	}

	fmt.Println(string(jsonOutput))
}
```

### Fetching iOS Updates

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/sofa"
)

func main() {
	fmt.Println("Fetching iOS updates...")

	data, err := sofa.GetiOSUpdates()
	if err != nil {
		log.Fatalf("Error fetching iOS updates: %v", err)
	}

	// Access structured data
	fmt.Printf("Latest iOS: %s (Build %s)\n", 
		data.OSVersions[0].Latest.ProductVersion, 
		data.OSVersions[0].Latest.Build)

	// Or convert to JSON for display/storage
	jsonOutput, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON output: %v", err)
	}

	fmt.Println(string(jsonOutput))
}
```

## Key Features of SOFA

### Machine-Readable Feed, RSS Feed, and Web UI

- **JSON Feed**: Provides detailed, machine-readable data optimized for automated tools and scripts

### Use Cases

- **XProtect Monitoring**: Keep track of the latest XProtect updates centrally to verify compliance with CIS/mSCP standards
- **Security Overviews**: Surface information on vulnerabilities (CVEs) and their exploitation status (KEV)
- **Track Countdowns**: Know both timestamp and days since a release was posted
- **Documentation Access**: Quick links to relevant Apple documentation and CVE information
- **Download Universal Mac Assistant**: Access the latest and all 'active' (currently signed) IPSW/Universal Mac Assistant (UMA) download links

## Available Data

### macOS Data

- OS versions and builds
- Security releases and updates
- XProtect payload information
- Supported device models
- CVE information including actively exploited vulnerabilities
- Installation app details (Universal Mac Assistant, IPSW)

### iOS Data

- OS versions and builds
- Security releases and updates
- Supported device information
- CVE details including actively exploited vulnerabilities

## Important Notes

SOFA has issued a deprecation notice for their feed URLs. Please ensure you're using the current endpoints:

- macOS feed: `https://sofafeed.macadmins.io/v1/macos_data_feed.json`
- iOS feed: `https://sofafeed.macadmins.io/v1/ios_data_feed.json`

The SDK already uses these updated endpoints.

## Examples

See the [examples](./examples) directory for complete working examples:

- [GetMacOSUpdates](./examples/GetMacOSUpdates/GetMacOSUpdates.go)
- [GetiOSUpdates](./examples/GetiOSUpdates/GetiOSUpdates.go)
- [ExportmacOSUpdatesToJson](./examples/ExportmacOSUpdatesToJson/ExportmacOSUpdatesToJson.go)
- [ExportiOSUpdatesToJson](./examples/ExportiOSUpdatesToJson/ExportiOSUpdatesToJson.go)

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the [MIT License](LICENSE).

## Links

- [SOFA Web UI](https://sofa.macadmins.io/)
- [SOFA GitHub Repository](https://github.com/macadmins/sofa)
