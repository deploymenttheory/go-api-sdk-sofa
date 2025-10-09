package client

import (
	"context"
	"fmt"

	"resty.dev/v3"
)

// Get executes a GET request following Resty v3 best practices.
// Simplified for SOFA API which only needs GET requests.
func (c *Client) Get(ctx context.Context, path string, result any) error {
	req := c.httpClient.R().
		SetContext(ctx).
		SetResult(result)

	return c.executeRequest(req, "GET", path)
}

// executeRequest is a centralized request executor that handles error processing.
func (c *Client) executeRequest(req *resty.Request, method, path string) error {
	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = req.Get(path)
	default:
		return fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return c.errorHandler.HandleError(resp)
	}

	return nil
}
