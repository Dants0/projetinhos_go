package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <url>")
		return
	}

	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Status não OK, %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler HTML: %v", err)
	}

	title := doc.Find("title").Text()
	description, _ := doc.Find("meta[name='description']").Attr("content")
	fmt.Println("Titulo da página: ", title)
	fmt.Println("Descrição da página: ", description)

	var links []string
	var removeDuplicates []string
	fmt.Println("\nLinks encontrados:")
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, href)
			removeDuplicates = RemoveDuplicates(links)
		}
	})
	fmt.Printf("- %s\n", removeDuplicates)
}