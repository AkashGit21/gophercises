package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/AkashGit21/gophercises/choose-your-own-adventure/arc"
)

func main() {
	port := flag.Int("port", 8082, "the port to start CYOA web application on")
	file := flag.String("file", "gophers.json", "the JSON file from which the stories are parsed")
	flag.Parse()

	chapters, err := arc.InitiateChapters(*file)
	if err != nil {
		log.Fatal(err)
	}

	h := arc.NewHandler(chapters)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
