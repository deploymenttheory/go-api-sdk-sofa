package client

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"resty.dev/v3"
)

// APIError represents an error from the SOFA API
type APIError struct {
	StatusCode int
	Message    string
	URL        string
	Method     string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}

// ErrorHandler centralizes error handling for all API requests
type ErrorHandler struct {
	logger *zap.Logger
}

// NewErrorHandler creates a new error handler
func NewErrorHandler(logger *zap.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

// HandleError processes API error responses and returns structured errors
func (eh *ErrorHandler) HandleError(resp *resty.Response) error {
	statusCode := resp.StatusCode()

	eh.logger.Error("API request failed",
		zap.Int("status_code", statusCode),
		zap.String("url", resp.Request.URL),
		zap.String("method", resp.Request.Method),
		zap.String("response_body", resp.String()),
	)

	return &APIError{
		StatusCode: statusCode,
		Message:    fmt.Sprintf("HTTP %d: %s", statusCode, http.StatusText(statusCode)),
		URL:        resp.Request.URL,
		Method:     resp.Request.Method,
	}
}

// Common error types
var (
	ErrInvalidResponse = fmt.Errorf("invalid response format")
	ErrRequestFailed   = fmt.Errorf("request failed")
)
