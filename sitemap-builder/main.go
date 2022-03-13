package main

import (
	"log"
	"net/http"
)

const (
	windowsAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246"
	TARGET_URL   = "TARGET URL"
	DEPTH        = 1
)

func main() {
	err := bfs(TARGET_URL, DEPTH)
	if err != nil {
		log.Fatal(err)
	}

	displayLinks()
}

// Perform a GET request and returns the response obtained
func fetchResponse(url string) (*http.Response, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", windowsAgent)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
