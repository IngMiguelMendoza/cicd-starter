package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	goodHeader := make(http.Header, 1)
	goodHeader.Add("Authorization", "ApiKey catsAreAwesome")

	headerMissingValue := make(http.Header)
	headerMissingValue.Add("Authorization", "catsaredog")

	headerMissingAuthKey := make(http.Header)

	tests := []struct {
		name        string
		input       http.Header
		expected    string
		expectedErr error
	}{
		{
			name:        "Valid header",
			input:       goodHeader,
			expected:    "catsAreAwesome",
			expectedErr: nil,
		},
		{
			name:        "Missing Authorization header",
			input:       headerMissingAuthKey,
			expected:    "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetAPIKey(tc.input)

			// Check if the result matches the expected value
			if result != tc.expected {
				t.Errorf("expected result: %s, got: %s", tc.expected, result)
			}

			// Check if the error matches the expected error
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
