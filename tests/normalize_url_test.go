package tests

import (
	"testing"

	"github.com/datsun80zx/webscrpr.git/internal"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "lowercase hostname",
			inputURL: "https://BLOG.Boot.Dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "handle empty path",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev",
		},
		{
			name:     "handle root path with trailing slash",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "preserve path case",
			inputURL: "https://blog.boot.dev/Path/To/Resource",
			expected: "blog.boot.dev/Path/To/Resource",
		},
		{
			name:     "handle deeper paths",
			inputURL: "https://blog.boot.dev/path/to/resource",
			expected: "blog.boot.dev/path/to/resource",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := internal.NormalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
