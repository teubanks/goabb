package chimux

import (
	"net/http"

	"github.com/go-chi/chi"
)

// ChiMux is a wrapper around Chi that helps it conform to the interface for GOA
type ChiMux struct {
	chi *chi.Mux
}

func New() *ChiMux {
	r := chi.NewRouter()
	return &ChiMux{chi: r}
}

func (c *ChiMux) Handle(method string, pattern string, handler http.HandlerFunc) {
	// Ensure the pattern handles both bare and trailing / patterns
	c.chi.Method(method, pattern, handler)
}

func (c *ChiMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.chi.ServeHTTP(w, r)
}
func (c *ChiMux) Vars(r *http.Request) map[string]string {
	vars := make(map[string]string)
	ctx := chi.RouteContext(r.Context())
	for i, k := range ctx.URLParams.Keys {
		vars[k] = ctx.URLParams.Values[i]
	}
	return vars
}
