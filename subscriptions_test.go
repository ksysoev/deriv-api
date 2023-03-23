package deriv

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestParseSubscription_ValidInput(t *testing.T) {
	input := `{"subscription": {"id": "123"}}`
	expected := SubscriptionResponse{Subscription: SubscriptionIDResponse{ID: "123"}}
	result, err := parseSubsciption(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, result)
	}
}

func TestParseSubscription_InvalidJSONInput(t *testing.T) {
	input := `{"subscription": {"id": "123", "status": "active"`
	_, err := parseSubsciption(input)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestParseSubscription_InvalidSubscriptionData(t *testing.T) {
	input := `{"subscription": {"id": "123", "status": "active"}, "error": {"code": "invalid_subscription"}}`
	expectedErr := &APIError{Code: "invalid_subscription"}
	_, err := parseSubsciption(input)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("Expected %+v, but got %+v", expectedErr, err)
	}
}

func TestParseSubscription_EmptyInput(t *testing.T) {
	_, err := parseSubsciption("")
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestParseSubscription_EmptySubscriptionData(t *testing.T) {
	input := `{}`
	expectedErr := fmt.Errorf("subscription ID is empty")
	_, err := parseSubsciption(input)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	if errors.Is(err, expectedErr) {
		t.Errorf("Expected %+v, but got %+v", expectedErr, err)
	}
}
