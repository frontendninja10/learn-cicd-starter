package auth

import (
	"errors"
	"net/http"
	"testing"
)


func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name			string
		header			http.Header
		expectedError	error
		errorMessage	string
	}{
		{
			name: "no header",
			header: http.Header{},
			expectedError: ErrNoAuthHeaderIncluded,
			errorMessage: "no authorization header included",
		},
		{
			name: "missing key",
			header: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedError: errors.New("malformed authorization header"),
			errorMessage: "malformed authorization header",
		},
		{
			name: "missing key name",
			header: http.Header{
				"Authorization": []string{"abc123"},
			},
			expectedError: errors.New("malformed authorization header"),
			errorMessage: "malformed authorization header",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GetAPIKey(tc.header)
			if err == nil || err.Error() != tc.expectedError.Error() {
				t.Errorf("expected: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}