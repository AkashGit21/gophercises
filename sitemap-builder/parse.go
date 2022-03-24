package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href  string
	Title string
}

var (
	stck []*html.Node
)

// Parse the HTML response and return links present within the same
func parseResponse(in *http.Response, url string) ([]string, error) {
	// storeResponse(in, url+".html")
	if in == nil {
		return nil, fmt.Errorf("no response found!")
	}
	defer in.Body.Close()

	var links []string
	if in.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(in.Body)
		if err != nil {
			log.Fatal(err)
		}

		r := bytes.NewReader(bodyBytes)
		doc, err := html.Parse(r)
		if err != nil {
			return nil, err
		}

		links = extractLinks(doc)
	} else {
		return nil, fmt.Errorf("Status not ok!")
	}

	return links, nil
}

func storeResponse(in *http.Response, fname string) error {
	if in == nil {
		return fmt.Errorf("No response found!")
	}
	defer in.Body.Close()

	if in.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(in.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		f, err := os.Create(fname)
		if err != nil {
			return err
		}

		f.WriteString(bodyString)
		fmt.Println("Successfully entered the data!")
	}
	return nil
}

// Extract the links from this node
func extractLinks(root *html.Node) []string {
	anchorChan := make(chan *html.Node)
	go dfs(root, 3, "a", anchorChan)

	res := make([]string, 0)
	for link := <-anchorChan; link != nil; link = <-anchorChan {
		// fmt.Println("\nLink found!")
		for _, attr := range link.Attr {
			if attr.Key == "href" {
				res = append(res, adjustLink(attr.Val))
			}
		}
	}
	return res
}

// Traverse through every sub-element of the node and pass the required childnodes using channel
func dfs(node *html.Node, nodeType int, tag string, results chan *html.Node) {
	defer close(results)
	if node != nil {
		stck = nil
		stck = append(stck, node)
	} else {
		return
	}

	for curSz := len(stck); curSz > 0; {
		ele := stck[curSz-1]
		stck = stck[:curSz-1]

		for c := ele.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ErrorNode {
				log.Printf("Found error node!")
				return
			} else if c.Type == html.NodeType(nodeType) && c.Data == tag {
				results <- c
			} else if tag == "*" && c.Type == html.NodeType(nodeType) {
				results <- c
			} else {
				stck = append(stck, c)
			}
		}
		curSz = len(stck)
	}
}

// Fix / Adjust the link to make it absolute
func adjustLink(link string) string {
	res := ""

	if strings.HasPrefix(link, "http") {
		if strings.HasPrefix(link, TARGET_URL) {
			// For absolute links already available
			return link
		}
		// Exclude all external links
	} else {
		// For referenced links
		link = strings.Trim(link, "/")
		link = strings.Split(link, "#")[0]
		link = strings.Split(link, "?")[0]
		res = TARGET_URL + link
	}

	return res
}
