package internal

import (
	"log"
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
