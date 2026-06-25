package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	goodHeader := http.Header{}
	goodHeader.Set("Authorization", "ApiKey abcd123")

	badHeader := http.Header{}
	badHeader.Set("Authorization", "abcd123")

	noHeader := http.Header{}

	tests := []struct {
		input     http.Header
		outputKey string
		haveErr   bool
	}{
		{input: goodHeader, outputKey: "abcd123", haveErr: false},
		{input: badHeader, outputKey: "", haveErr: true},
		{input: noHeader, outputKey: "", haveErr: true},
	}

	for _, tc := range tests {
		key, err := GetAPIKey(tc.input)

		if (err != nil) != tc.haveErr {
			t.Errorf("expected error to be: %v, got: %v", tc.haveErr, (err != nil))
		}

		if key != tc.outputKey {
			t.Errorf("expected key: %v, got: %v", tc.outputKey, key)
		}
	}

}
