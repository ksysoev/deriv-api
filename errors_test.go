package deriv

import (
	"encoding/json"
	"errors"
	"testing"
)

var expectedDetails json.RawMessage = []byte(`{"TestKey":"TestValue"}`)

func TestAPIError_Error(t *testing.T) {
	err := &APIError{
		Code:    "test-code",
		Message: "test-message",
		Details: &expectedDetails,
	}

	expected := err.Message
	actual := err.Error()

	if actual != expected {
		t.Errorf("Error() returned %q, expected %q", actual, expected)
	}
}

func TestParseError_ValidResponse(t *testing.T) {
	errorResponse := apiErrorResponse{
		Error: APIError{
			Code:    "test-code",
			Message: "test-message",
			Details: &expectedDetails,
		},
	}

	rawResponse, err := json.Marshal(&errorResponse)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := &errorResponse.Error
	actual := parseError(rawResponse)

	if actual.Error() != expected.Error() {
		t.Errorf("parseError() returned %v, expected %v", actual, expected)
	}
}

func TestParseError_InvalidResponse(t *testing.T) {
	rawResponse := []byte("invalid-json")

	expected := errors.New("invalid character 'i' looking for beginning of value")
	actual := parseError(rawResponse)

	if actual.Error() != expected.Error() {
		t.Errorf("parseError() returned %v, expected %v", actual, expected)
	}
}

func TestParseError_EmptyErrorResponse(t *testing.T) {
	rawResponse := []byte("{}")

	actual := parseError(rawResponse)

	if actual != nil {
		t.Errorf("parseError() returned %v, expected %v", actual, nil)
	}
}

func TestParseError_EmptyAPIError(t *testing.T) {
	errorResponse := apiErrorResponse{
		Error: APIError{},
	}

	rawResponse, err := json.Marshal(&errorResponse)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	actual := parseError(rawResponse)

	if actual != nil {
		t.Errorf("parseError() returned %v, expected %v", actual, nil)
	}
}
func TestAPIError_ParseDetails_ValidDetails(t *testing.T) {
	details := struct {
		TestKey string `json:"TestKey"`
	}{}

	apiErr := &APIError{
		Code:    "test-code",
		Message: "test-message",
		Details: &expectedDetails,
	}

	if err := apiErr.ParseDetails(&details); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if details.TestKey != "TestValue" {
		t.Errorf("ParseDetails() did not parse details correctly, expected %q, got %q", "test-value", details.TestKey)
	}
}

func TestAPIError_ParseDetails_EmptyDetails(t *testing.T) {
	details := struct {
		TestKey string `json:"TestKey"`
	}{}

	apiErr := &APIError{
		Code:    "test-code",
		Message: "test-message",
		Details: nil,
	}

	err := apiErr.ParseDetails(&details)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if details.TestKey != "" {
		t.Errorf("ParseDetails() did not handle empty details correctly, expected %q, got %q", "", details.TestKey)
	}
}
