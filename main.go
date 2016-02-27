package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(j int, s *goquery.Selection) {
			fmt.Print(s.Text(), "\t")
		})
		fmt.Println()
	})
}
