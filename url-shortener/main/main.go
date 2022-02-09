package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/AkashGit21/gophercises/url-shortener"
)

// var data = `
// a: a string from struct A
// b: a string from struct B
// `
var yaml = `
- path: /AkashGupta
  url: https://github.com/AkashGit21
- path: /urlshort-final
  url: github.com/AkashGit21/gophercises/url-shortener
`

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/gophercies": "https://www.github.com/AkashGit21/gophercises",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
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
	fmt.Fprintln(w, "404 Page not found!")
}
