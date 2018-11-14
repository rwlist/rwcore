package articles

import (
	"github.com/go-chi/chi"
	"github.com/rwlist/rwcore/app/auth"
)

func (m *Module) Router() chi.Router {
	z := resp{}

	r := chi.NewRouter()
	r.Use(auth.HasRole("articles"))

	// r.Post("/add", m.addOne)
	// r.Post("/addMany", m.addMany)
	r.Get("/all", z.getAll)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(fetchArticle)
		r.Get("/", z.get)
		r.Post("/click", z.onClick)
		r.Post("/read/status", z.setReadStatus)
		r.Post("/rating/change", z.changeRating)
		// r.Patch("/patch", m.patch)
	})

	return r
}
