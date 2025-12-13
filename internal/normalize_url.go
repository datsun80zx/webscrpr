package internal

import (
	"log"
	"net/url"
	"strings"
)

func NormalizeURL(inputURL string) (string, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("scheme: %v\n\n", u.Scheme)
	// fmt.Printf("user:%v\n\n", u.User)
	// fmt.Printf("host: %v\n\n", u.Host)
	// fmt.Printf("path: %v\n\n", u.Path)
	// fmt.Printf("rawpath: %v\n\n", u.RawPath)

	cleanHost := strings.TrimSuffix(strings.ToLower(u.Host), "/")
	cleanPath := strings.TrimSuffix(u.Path, "/")

	cleanURL := cleanHost + cleanPath
	return cleanURL, nil
}
