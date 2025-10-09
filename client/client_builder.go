package client

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// ClientBuilder provides a fluent interface for building SOFA clients
type ClientBuilder struct {
	baseURL    string
	timeout    time.Duration
	retryCount int
	retryWait  time.Duration
	userAgent  string
	debug      bool
	logger     *zap.Logger
}

// NewClientBuilder creates a new client builder with default values
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		baseURL:    BaseURLV1,
		timeout:    30 * time.Second,
		retryCount: 3,
		retryWait:  1 * time.Second,
		userAgent:  "go-api-sdk-sofa/1.0.0",
		debug:      false,
	}
}

// WithBaseURL sets the base URL for the API (default: v1)
func (cb *ClientBuilder) WithBaseURL(baseURL string) *ClientBuilder {
	cb.baseURL = baseURL
	return cb
}

// WithV1 sets the base URL to the v1 API endpoint
func (cb *ClientBuilder) WithV1() *ClientBuilder {
	cb.baseURL = BaseURLV1
	return cb
}

// WithV2 sets the base URL to the v2 API endpoint
func (cb *ClientBuilder) WithV2() *ClientBuilder {
	cb.baseURL = BaseURLV2
	return cb
}

// WithTimeout sets the request timeout (default: 30 seconds)
func (cb *ClientBuilder) WithTimeout(timeout time.Duration) *ClientBuilder {
	cb.timeout = timeout
	return cb
}

// WithRetry configures retry settings (default: 3 retries, 1 second wait)
func (cb *ClientBuilder) WithRetry(count int, wait time.Duration) *ClientBuilder {
	cb.retryCount = count
	cb.retryWait = wait
	return cb
}

// WithUserAgent sets the user agent string (default: go-api-sdk-sofa/1.0.0)
func (cb *ClientBuilder) WithUserAgent(userAgent string) *ClientBuilder {
	cb.userAgent = userAgent
	return cb
}

// WithDebug enables or disables debug mode (default: false)
func (cb *ClientBuilder) WithDebug(debug bool) *ClientBuilder {
	cb.debug = debug
	return cb
}

// WithLogger sets a custom logger (default: no-op logger)
func (cb *ClientBuilder) WithLogger(logger *zap.Logger) *ClientBuilder {
	cb.logger = logger
	return cb
}

// Build creates and returns the configured SOFA client
func (cb *ClientBuilder) Build() (*Client, error) {
	if cb.logger == nil {
		if cb.debug {
			cb.logger, _ = zap.NewDevelopment()
		} else {
			cb.logger = zap.NewNop()
		}
	}

	if cb.baseURL == "" {
		return nil, fmt.Errorf("base URL is required")
	}
	if cb.timeout <= 0 {
		return nil, fmt.Errorf("timeout must be positive")
	}
	if cb.retryCount < 0 {
		return nil, fmt.Errorf("retry count cannot be negative")
	}

	config := Config{
		BaseURL:    cb.baseURL,
		Logger:     cb.logger,
		Timeout:    cb.timeout,
		RetryCount: cb.retryCount,
		RetryWait:  cb.retryWait,
		UserAgent:  cb.userAgent,
		Debug:      cb.debug,
	}

	return NewClient(config)
}

// MustBuild creates the SOFA client and panics if there's an error
func (cb *ClientBuilder) MustBuild() *Client {
	client, err := cb.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build SOFA client: %v", err))
	}
	return client
}

// Validate checks the current configuration without building the client
func (cb *ClientBuilder) Validate() error {
	if cb.baseURL == "" {
		return fmt.Errorf("base URL is required")
	}
	if cb.timeout <= 0 {
		return fmt.Errorf("timeout must be positive")
	}
	if cb.retryCount < 0 {
		return fmt.Errorf("retry count cannot be negative")
	}
	return nil
}

// Clone creates a copy of the client builder
func (cb *ClientBuilder) Clone() *ClientBuilder {
	return &ClientBuilder{
		baseURL:    cb.baseURL,
		timeout:    cb.timeout,
		retryCount: cb.retryCount,
		retryWait:  cb.retryWait,
		userAgent:  cb.userAgent,
		debug:      cb.debug,
		logger:     cb.logger,
	}
}

// Public convenience functions for common client creation patterns

// NewDefaultClient creates a SOFA client with default v1 settings
func NewDefaultClient() (*Client, error) {
	return NewClientBuilder().Build()
}

// NewV1Client creates a SOFA client configured for v1 API
func NewV1Client() (*Client, error) {
	return NewClientBuilder().WithV1().Build()
}

// NewV2Client creates a SOFA client configured for v2 API
func NewV2Client() (*Client, error) {
	return NewClientBuilder().WithV2().Build()
}

// NewClientWithOptions creates a SOFA client with custom options
func NewClientWithOptions(debug bool, timeout time.Duration, userAgent string) (*Client, error) {
	builder := NewClientBuilder().
		WithDebug(debug).
		WithUserAgent(userAgent)

	if timeout > 0 {
		builder = builder.WithTimeout(timeout)
	}

	return builder.Build()
}
