package deriv

import (
	"encoding/json"
	"fmt"
)

var (
	ErrConnectionClosed    = fmt.Errorf("connection closed")
	ErrEmptySubscriptionID = fmt.Errorf("subscription ID is empty")
	ErrInvalidSchema       = fmt.Errorf("invalid endpoint scheme")
	ErrInvalidAppID        = fmt.Errorf("invalid app ID")
	ErrInvalidLanguage     = fmt.Errorf("invalid language")
)

// APIError represents an error returned by the Deriv API service.
type APIError struct {
	Details *json.RawMessage `json:"details"`
	Code    string           `json:"code"`
	Message string           `json:"message"`
}

// Error returns the error message associated with the APIError.
func (e *APIError) Error() string {
	return e.Message
}

// ParseDetails parses the details field of the APIError into the provided value.
func (e *APIError) ParseDetails(v any) error {
	if e.Details == nil {
		return nil
	}

	return json.Unmarshal(*e.Details, v)
}

// apiErrorResponse represents an error response returned by the Deriv API service.
type apiErrorResponse struct {
	// Error is the APIError associated with the response.
	Error APIError `json:"error"`
}

// parseError parses a JSON error response from the Deriv API service into an error.
// If the response is not a valid JSON-encoded apiErrorResponse, an error is returned.
// If the apiErrorResponse contains a non-empty APIError, it is returned as an error.
// Otherwise, nil is returned.
func parseError(rawResponse []byte) error {
	var errorResponse apiErrorResponse

	err := json.Unmarshal(rawResponse, &errorResponse)
	if err != nil {
		return err
	}

	if errorResponse.Error.Code != "" {
		return &errorResponse.Error
	}

	return nil
}
