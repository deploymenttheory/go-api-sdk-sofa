package sofa

import (
	"context"

	"github.com/deploymenttheory/go-api-sdk-sofa/client"
	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/models"
)

// GetMacOSSofaFeedV1 fetches macOS update data from the v1 API
func (c *Client) GetMacOSSofaFeedV1(ctx context.Context) (*models.MacOSV1FeedResponse, error) {
	var result models.MacOSV1FeedResponse
	if err := c.client.Get(ctx, client.EndpointMacOSV1, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetiOSSofaFeedV1 fetches iOS update data from the v1 API
func (c *Client) GetiOSSofaFeedV1(ctx context.Context) (*models.IOSV1FeedResponse, error) {
	var result models.IOSV1FeedResponse
	if err := c.client.Get(ctx, client.EndpointIOSV1, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
