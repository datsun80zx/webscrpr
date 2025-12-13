package main

import (
	"fmt"
	"log"

	"github.com/datsun80zx/webscrpr.git/internal"
)

func main() {
	url := "https://www.boot.dev/lessons/98ac1f38-22dd-4682-b114-8638a0625567"

	parsedURL, err := internal.NormalizeURL(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("func returns:", parsedURL)
}
