package auth

import (
	"net/http"
	"testing"
)


func TestGetAPIKey_NoHeader(t *testing.T) {
	header := http.Header{}
	_, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("expected error %v, got nil", ErrNoAuthHeaderIncluded)
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
}