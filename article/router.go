package article

import (
	"github.com/go-chi/chi"
	"github.com/rwlist/rwcore/auth"
)

type Router struct{ *chi.Mux }

func NewRouter(c *Controller, auth *auth.Middleware, mid *Middleware) Router {
	r := chi.NewRouter()
	r.Use(auth.HasRole("article"))

	// r.Post("/add", m.addOne)
	// r.Post("/addMany", m.addMany)
	r.Get("/all", c.getAll)
	r.Post("/addURL", c.addURL)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(mid.UpdateContext)
		r.Get("/", c.get)
		r.Post("/click", c.onClick)
		r.Post("/read/status", c.setReadStatus)
		r.Post("/rating/change", c.changeRating)
		r.Post("/tags/remove", c.removeTag)
		r.Post("/tags/add", c.addTag)
		// r.Patch("/patch", m.patch)
	})
	return Router{r}
}
