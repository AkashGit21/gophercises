package arc

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var defaultHandlerTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Choose Your Own Adventure</title>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{range .Paragraphs}}
    <p>{{.}}</p>
  {{end}}
  <ul>
    {{range .Options}}
      <li> <a href="/{{.Arc}}">{{.Description}}</a> </li>
    {{end}}
  </ul>
</body>
</html>`
var tpl *template.Template

type Story map[string]Chapter

type Chapter struct {
	Title      string    `json:"title,omitempty"`
	Paragraphs []string  `json:"story,omitempty"`
	Options    []*Option `json:"options,omitempty"`
}

type Option struct {
	Description string `json:"text,omitempty"`
	Arc         string `josn:"arc,omitempty"`
}

type handler struct {
	s Story
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

func InitiateChapters(JSONInput string) (Story, error) {

	data, err := os.ReadFile(JSONInput)
	if err != nil {
		return nil, err
	}

	var story Story

	err = json.Unmarshal(data, &story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	arc := r.URL.Path
	arc = strings.TrimPrefix(arc, "/")
	if arc == "" {
		arc = "intro"
	}

	if chapter, ok := h.s[arc]; !ok {
		http.Error(rw, fmt.Sprintf("Arc %s not found!", arc), http.StatusNotFound)
	} else {

		err := tpl.Execute(rw, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(rw, "Something went wrong...", http.StatusInternalServerError)
		}
	}
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}
