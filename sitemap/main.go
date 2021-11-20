package main

import (
	"flag"
	"io"
	"net/http"
	"os"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url to you want to build a sitemap for")
	flag.Parse()

	// fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

/*
	1. GET the webpage
	2. Parse the HTML
	3. Build proper URLs with our links
	4. Filter out any links with a different domain
	5. Find all pages (BFS Breadth First Search)
	6. Print out XML
*/
