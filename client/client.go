package client

import (
	"time"

	"go.uber.org/zap"
	"resty.dev/v3"
)

// Client represents the main SOFA API client
type Client struct {
	httpClient   *resty.Client
	logger       *zap.Logger
	errorHandler *ErrorHandler
	baseURL      string
}

// Config holds configuration for the client
type Config struct {
	BaseURL    string
	Logger     *zap.Logger
	Timeout    time.Duration
	RetryCount int
	RetryWait  time.Duration
	UserAgent  string
	Debug      bool
}

// NewClient creates a new SOFA API client
func NewClient(config Config) (*Client, error) {
	if config.BaseURL == "" {
		config.BaseURL = BaseURLV1
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}
	if config.RetryCount == 0 {
		config.RetryCount = 3
	}
	if config.RetryWait == 0 {
		config.RetryWait = 1 * time.Second
	}
	if config.UserAgent == "" {
		config.UserAgent = "go-api-sdk-sofa/1.0.0"
	}
	if config.Logger == nil {
		config.Logger = zap.NewNop()
	}

	httpClient := resty.New()

	httpClient.
		SetBaseURL(config.BaseURL).
		SetTimeout(config.Timeout).
		SetRetryCount(config.RetryCount).
		SetRetryWaitTime(config.RetryWait).
		SetRetryMaxWaitTime(config.RetryWait*10).
		SetHeader("User-Agent", config.UserAgent).
		SetHeader("Accept", "application/json")

	if config.Debug {
		httpClient.SetDebug(true)
	}

	httpClient.AddRequestMiddleware(func(c *resty.Client, req *resty.Request) error {
		config.Logger.Info("API request",
			zap.String("method", req.Method),
			zap.String("url", req.URL),
		)
		return nil
	})

	httpClient.AddResponseMiddleware(func(c *resty.Client, resp *resty.Response) error {
		config.Logger.Info("API response",
			zap.String("method", resp.Request.Method),
			zap.String("url", resp.Request.URL),
			zap.Int("status_code", resp.StatusCode()),
			zap.String("status", resp.Status()),
		)
		return nil
	})

	client := &Client{
		httpClient:   httpClient,
		logger:       config.Logger,
		errorHandler: NewErrorHandler(config.Logger),
		baseURL:      config.BaseURL,
	}

	return client, nil
}

// Close closes the HTTP client and cleans up resources
func (c *Client) Close() error {
	if c.httpClient != nil {
		c.httpClient.Close()
	}
	return nil
}

// GetHTTPClient returns the underlying HTTP client for testing purposes
func (c *Client) GetHTTPClient() *resty.Client {
	return c.httpClient
}
