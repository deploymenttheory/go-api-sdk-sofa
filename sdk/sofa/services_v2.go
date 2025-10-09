package sofa

import (
	"context"

	"github.com/deploymenttheory/go-api-sdk-sofa/client"
	"github.com/deploymenttheory/go-api-sdk-sofa/sdk/models"
)

// GetMacOSSofaFeedV2 fetches macOS update data from the v2 API
func (c *Client) GetMacOSSofaFeedV2(ctx context.Context) (*models.MacOSV2FeedResponse, error) {
	var result models.MacOSV2FeedResponse
	if err := c.client.Get(ctx, client.EndpointMacOSV2, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetiOSSofaFeedV2 fetches iOS update data from the v2 API
func (c *Client) GetiOSSofaFeedV2(ctx context.Context) (*models.IOSFeedResponse, error) {
	var result models.IOSFeedResponse
	if err := c.client.Get(ctx, client.EndpointIOSV2, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTvOSSofaFeedV2 fetches tvOS update data from the v2 API
func (c *Client) GetTvOSSofaFeedV2(ctx context.Context) (*models.TVOSV2FeedResponse, error) {
	var result models.TVOSV2FeedResponse
	if err := c.client.Get(ctx, client.EndpointTVOSV2, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWatchOSSofaFeedV2 fetches watchOS update data from the v2 API
func (c *Client) GetWatchOSSofaFeedV2(ctx context.Context) (*models.WatchOSV2FeedResponse, error) {
	var result models.WatchOSV2FeedResponse
	if err := c.client.Get(ctx, client.EndpointWatchOSV2, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetVisionOSSofaFeedV2 fetches visionOS update data from the v2 API
func (c *Client) GetVisionOSSofaFeedV2(ctx context.Context) (*models.VisionOSV2FeedResponse, error) {
	var result models.VisionOSV2FeedResponse
	if err := c.client.Get(ctx, client.EndpointVisionOSV2, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSafariSofaFeedV2 fetches Safari update data from the v2 API
func (c *Client) GetSafariSofaFeedV2(ctx context.Context) (*models.SafariV2FeedResponse, error) {
	var result models.SafariV2FeedResponse
	if err := c.client.Get(ctx, client.EndpointSafariV2, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
