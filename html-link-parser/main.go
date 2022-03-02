package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var (
	links []Link
)

func main() {

	s := `
	<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
	`

	if err := parseHTML(s); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Links: %+v\n", links)
}

func parseHTML(str string) error {

	r := strings.NewReader(str)

	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	getLinks(doc)
	return nil
}

func getLinks(n *html.Node) (string, error) {

	res := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if value, err := getLinks(c); err != nil {
			return "", err
		} else {
			res += value
		}
	}

	switch n.Type {
	case html.CommentNode:
		return "", nil
	case html.TextNode:
		return n.Data, nil
	case html.ElementNode:
		if n.Data == "a" {

			href := ""
			for _, attribute := range n.Attr {
				if attribute.Key == "href" {
					href = attribute.Val
				}
			}

			res = strings.Join(strings.Fields(res), " ")
			// res = strings.ReplaceAll(res, "\n", " ")
			// res = strings.ReplaceAll(res, "\t", "")
			links = append(links, Link{
				Href: href,
				Text: res,
			})
		}
	}
	return res, nil
}
