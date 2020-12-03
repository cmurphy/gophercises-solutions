package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/cmurphy/gophercises/cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "JSON file containing the CYOA story sequence")
	flag.Parse()
	story, err := cyoa.ReadStory(*file)
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", cyoa.NewHandler(story)))
}
