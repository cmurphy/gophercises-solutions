package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gophercises/urlshort"
)

func main() {
	yamlFile := flag.String("yaml", "", "a YAML file mapping paths to URLs")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := []byte(`
   - path: /urlshort
     url: https://github.com/gophercises/urlshort
   - path: /urlshort-final
     url: https://github.com/gophercises/urlshort/tree/solution
`)
	if *yamlFile != "" {
		var err error
		yaml, err = ioutil.ReadFile(*yamlFile)
		if err != nil {
			panic(err)
		}
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
