package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
    <link rel="stylesheet" href="css/style.css">
  </head>

  <body>
    <h1>{{.Title}}</h1>

      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
        
        {{if .Options}}
          <ul>
            {{range .Options}}
              <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
          </ul>
        {{else}}
          <h3>The End</h3>
        {{end}}
  </body>
</html>
`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(rw, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(rw, "Something went wrong...", http.StatusInternalServerError)

		}
		return
	}
	http.Error(rw, "Chapter not found...", http.StatusNotFound)
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
