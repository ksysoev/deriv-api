package deriv

import (
	"testing"
)

func TestNewDerivAPI(t *testing.T) {
	// Valid endpoint and origin
	endpoint := "wss://example.com/ws"
	origin := "https://example.com"
	appID := 123
	lang := "en"
	api, err := NewDerivAPI(endpoint, appID, lang, origin)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if api.Endpoint.String() != endpoint+"?app_id=123&l=en" {
		t.Errorf("Unexpected endpoint: got %v, want %v", api.Endpoint.String(), endpoint)
	}
	if api.Origin.String() != origin {
		t.Errorf("Unexpected origin: got %v, want %v", api.Origin.String(), origin)
	}
	if api.AppID != appID {
		t.Errorf("Unexpected app ID: got %v, want %v", api.AppID, appID)
	}
	if api.Lang != lang {
		t.Errorf("Unexpected language: got %v, want %v", api.Lang, lang)
	}

	// Invalid endpoint scheme
	endpoint = "https://example.com/ws"
	origin = "https://example.com"
	appID = 123
	lang = "en"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid app ID
	endpoint = "wss://example.com/ws"
	origin = "https://example.com"
	appID = -1
	lang = "en"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid language
	endpoint = "wss://example.com/ws"
	origin = "https://example.com"
	appID = 123
	lang = "eng"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid endpoint URL
	endpoint = ":invalid:"
	origin = "https://example.com"
	appID = 123
	lang = "en"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid origin URL
	endpoint = "wss://example.com/ws"
	origin = ":invalid:"
	appID = 123
	lang = "en"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
