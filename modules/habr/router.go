package habr

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (m *Module) Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", m.index)
	r.Route("/{name}", func(r chi.Router) {
		r.Get("/data", m.data)
	})

	return r
}

func (m *Module) index(w http.ResponseWriter, r *http.Request) {

}

func (m *Module) data(w http.ResponseWriter, r *http.Request) {

}
