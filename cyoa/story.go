package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
)

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
	//t := template.Must(template.New("").Parse(defaultHandlerTmpl))
}

type handler struct {
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
