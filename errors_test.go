package deriv

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	err := &APIError{
		Code:    "test-code",
		Message: "test-message",
		Details: map[string]interface{}{"test-key": "test-value"},
	}

	expected := err.Message
	actual := err.Error()

	if actual != expected {
		t.Errorf("Error() returned %q, expected %q", actual, expected)
	}
}

func TestParseError_ValidResponse(t *testing.T) {
	errorResponse := APIErrorResponse{
		Error: APIError{
			Code:    "test-code",
			Message: "test-message",
			Details: map[string]interface{}{"test-key": "test-value"},
		},
	}

	rawResponse, err := json.Marshal(&errorResponse)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &errorResponse.Error
	actual := parseError(string(rawResponse))

	if actual.Error() != expected.Error() {
		t.Errorf("parseError() returned %v, expected %v", actual, expected)
	}
}

func TestParseError_InvalidResponse(t *testing.T) {
	rawResponse := "invalid-json"

	expected := errors.New("invalid character 'i' looking for beginning of value")
	actual := parseError(rawResponse)

	if actual.Error() != expected.Error() {
		t.Errorf("parseError() returned %v, expected %v", actual, expected)
	}
}

func TestParseError_EmptyErrorResponse(t *testing.T) {
	rawResponse := "{}"

	expected := (error)(nil)
	actual := parseError(rawResponse)

	if actual != expected {
		t.Errorf("parseError() returned %v, expected %v", actual, expected)
	}
}

func TestParseError_EmptyAPIError(t *testing.T) {
	errorResponse := APIErrorResponse{
		Error: APIError{},
	}

	rawResponse, err := json.Marshal(&errorResponse)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := (error)(nil)
	actual := parseError(string(rawResponse))

	if actual != expected {
		t.Errorf("parseError() returned %v, expected %v", actual, expected)
	}
}
