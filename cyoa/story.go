package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
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
              <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
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
	err := tpl.Execute(rw, h.s["intro"])
	if err != nil {
		panic(err)
	}
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
	Chapter string `json:"chapter"`
}
