package deriv

import (
	"encoding/json"
	"fmt"
)

var (
	ErrConnectionClosed = fmt.Errorf("connection closed")
)

// APIError represents an error returned by the Deriv API service.
type APIError struct {
	Details map[string]interface{} `json:"details"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
}

// Error returns the error message associated with the APIError.
func (e *APIError) Error() string {
	return e.Message
}

// APIErrorResponse represents an error response returned by the Deriv API service.
type APIErrorResponse struct {
	// Error is the APIError associated with the response.
	Error APIError `json:"error"`
}

// parseError parses a JSON error response from the Deriv API service into an error.
// If the response is not a valid JSON-encoded APIErrorResponse, an error is returned.
// If the APIErrorResponse contains a non-empty APIError, it is returned as an error.
// Otherwise, nil is returned.
func parseError(rawResponse []byte) error {
	var errorResponse APIErrorResponse

	err := json.Unmarshal(rawResponse, &errorResponse)
	if err != nil {
		return err
	}

	if errorResponse.Error.Code != "" {
		return &errorResponse.Error
	}

	return nil
}
