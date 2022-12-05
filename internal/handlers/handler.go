package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-axesthump-adventure/internal/arc"
	"html/template"
	"net/http"
)

const mainPage = "intro"
const arcTitle = "arcTitle"

var tmpl = template.Must(
	template.New("data").Parse(
		`
<h1>{{ .Title}}</h1>
{{range .Story}}
     <p>{{.}}</p>
{{end}}
<p></p>
{{range .Options}}
     <p><a href="{{.ArcName}}">{{.Text}}</a></p>
{{end}}
`,
	),
)

var errorTmpl = template.Must(
	template.New("data").Parse(
		`
<h1>Error!</h1>
<h2>{{ .}}</h2>
`,
	),
)

type AppHandler struct {
	Router chi.Router
	Arcs   map[string]arc.Arc
}

func newRouter(h *AppHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "text/html"))

	r.Get("/", h.openArc)
	r.Get("/{arcTitle}", h.openArc)
	return r
}

func NewAppHandler() (*AppHandler, error) {
	arcs, err := arc.GetArcs()
	if err != nil {
		return nil, err
	}
	handler := &AppHandler{
		Arcs: arcs,
	}
	handler.Router = newRouter(handler)
	return handler, nil
}

func (h *AppHandler) openArc(w http.ResponseWriter, r *http.Request) {
	arcID := chi.URLParam(r, arcTitle)
	if len(arcID) == 0 {
		arcID = mainPage
	}
	err := tmpl.Execute(w, h.Arcs[arcID])
	if err != nil {
		err = errorTmpl.Execute(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
