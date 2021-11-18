package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lyx0/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 8080, "the port number to run the webserver on")
	filename := flag.String("file", "gopher.json", "the JSON story with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story, nil)
	fmt.Printf("Starting the server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
