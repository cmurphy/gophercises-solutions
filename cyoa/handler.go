package cyoa

import (
	"html/template"
	"net/http"
	"strings"
)

const pageTemplate = `
<h1>{{ .Title }}</h1>
{{range .Paragraphs}}
<p>{{.}}</p>
{{end}}
<ul>
{{range .Options}}
<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
{{end}}
</ul>
`

type storyHandler struct {
	story Story
}

// NewHandler returns a story handler
func NewHandler(s Story) http.Handler {
	return &storyHandler{story: s}
}

// ServeHTTP serves the app
func (s *storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}
	page := strings.TrimPrefix(path, "/")
	arc, ok := s.story[page]
	if !ok {
		http.NotFound(w, r)
		return
	}
	t, err := template.New("page.html").Parse(pageTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, &arc)
}
