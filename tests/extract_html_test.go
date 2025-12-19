package tests

import (
	"net/url"
	"reflect"
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

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><a href="https://blog.boot.dev/articles">Articles</a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/articles"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><a href="/about">About</a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/about"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLMultipleLinks(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body>
		<a href="/home">Home</a>
		<a href="/about">About</a>
		<a href="https://example.com/external">External</a>
	</body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://blog.boot.dev/home",
		"https://blog.boot.dev/about",
		"https://example.com/external",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLNoLinks(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><p>No links here</p></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLNestedContent(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><a href="/posts"><span><strong>Blog Posts</strong></span></a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/posts"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><img src="https://blog.boot.dev/logo.png" alt="Logo"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><img src="/images/logo.png" alt="Logo"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/images/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLMultiple(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body>
		<img src="/header.png" alt="Header">
		<img src="/content.jpg" alt="Content">
		<img src="https://cdn.example.com/external.gif" alt="External">
	</body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://blog.boot.dev/header.png",
		"https://blog.boot.dev/content.jpg",
		"https://cdn.example.com/external.gif",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLNoImages(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><p>No images here</p></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLMissingAlt(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><img src="/logo.png"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLRelativeNoDash(t *testing.T) {
	inputURL := "https://blog.boot.dev/posts/"
	inputBody := `<html><body><img src="thumbnail.png" alt="Thumb"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := internal.GetImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/posts/thumbnail.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
