package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://example.com/"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("h2").Each(func(index int, element *goquery.Selection) {
		title := element.Text()
		fmt.Printf("Title %d: %s\n", index+1, title)
	})

	document.Find("p").Each(func(index int, element *goquery.Selection) {
		paragraph := element.Text()
		fmt.Printf("Paragraph %d: %s\n", index+1, paragraph)
	})
}
