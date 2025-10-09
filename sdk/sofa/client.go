package sofa

import (
	"github.com/deploymenttheory/go-api-sdk-sofa/client"
)

// Client provides access to SOFA API Services
type Client struct {
	client *client.Client
}

// NewClient creates a new SOFA services client
func NewClient(c *client.Client) *Client {
	return &Client{
		client: c,
	}
}
