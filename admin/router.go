package admin

import (
	"github.com/go-chi/chi"
	"github.com/rwlist/rwcore/auth"
)

type Router struct{ *chi.Mux }

func NewRouter(c *Controller, middleware *auth.Middleware) Router {
	r := chi.NewRouter()
	r.Use(middleware.HasRole("admin"))
	r.Get("/test", c.test)
	return Router{r}
}
