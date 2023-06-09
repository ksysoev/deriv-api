package deriv

import (
	"encoding/json"
)

// APIError represents an error returned by the Deriv API service.
type APIError struct {
	// Code is a string representing the error code returned by the Deriv API service.
	Code string `json:"code"`
	// Message is a human-readable string describing the error.
	Message string `json:"message"`
	// Details is a map of additional error details.
	Details map[string]interface{} `json:"details"`
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
func parseError(rawResponse string) error {
	var errorResponse APIErrorResponse

	err := json.Unmarshal([]byte(rawResponse), &errorResponse)
	if err != nil {
		return err
	}

	if errorResponse.Error.Code != "" {
		return &errorResponse.Error
	}

	return nil
}
