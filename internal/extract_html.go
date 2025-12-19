package internal

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	h1 := doc.Find("h1")
	return h1.Text()
}

func GetFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	p := doc.Find("main p").First()
	if p.Length() == 0 {
		p = doc.Find("p").First()
	}
	return p.Text()
}

func GetURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("there was an error creating the Document object: %w", err)
	}

	URLList := []string{}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		parsedHref, err := url.Parse(href)
		if err != nil {
			return
		}
		absolute := baseURL.ResolveReference(parsedHref)
		URLList = append(URLList, absolute.String())
	})

	return URLList, nil
}

func GetImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("there was an error creating the Document object: %w", err)
	}

	imgList := []string{}

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists {
			return
		}

		parsedSrc, err := url.Parse(src)
		if err != nil {
			return
		}

		absolute := baseURL.ResolveReference(parsedSrc)
		imgList = append(imgList, absolute.String())
	})

	return imgList, nil
}
