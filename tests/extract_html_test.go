package tests

import (
	"testing"

	"github.com/datsun80zx/webscrpr.git/internal"
)

func TestGetH1FromHTML(t *testing.T) {
	h1Tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "normalHTML",
			inputHTML: "<html><body><h1>Welcome to Boot.dev</h1><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "Welcome to Boot.dev",
		},
		{
			name:      "missingH1",
			inputHTML: "<html><body><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "",
		},
		{
			name:      "H1 diff location",
			inputHTML: "<html><h1>Welcome to Boot.dev</h1><body><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "Welcome to Boot.dev",
		},
	}
	for i, tc := range h1Tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := internal.GetH1FromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected H1: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	paragraphTests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "normalHTML",
			inputHTML: "<html><body><h1>Welcome to Boot.dev</h1><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "Learn to code by building real projects.",
		},
		{
			name:      "missing paragraph",
			inputHTML: "<html><body><main></main></body></html>",
			expected:  "",
		},
		{
			name:      "paragraph diff location",
			inputHTML: "<html><h1>Welcome to Boot.dev</h1><body><p>this isn't the paragraph I want</p><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "Learn to code by building real projects.",
		},
		{
			name:      "no main tag",
			inputHTML: "<html><body><h1>Welcome to Boot.dev</h1><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></body></html>",
			expected:  "Learn to code by building real projects.",
		},
	}
	for i, tc := range paragraphTests {
		t.Run(tc.name, func(t *testing.T) {
			actual := internal.GetFirstParagraphFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected paragraph: %v\nactual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
