package admin

import (
	"github.com/go-chi/chi"
	"github.com/rwlist/rwcore/auth"
)

type Router *chi.Mux

func NewRouter(c Controller) Router {
	r := chi.NewRouter()
	r.Use(auth.HasRole("admin"))
	r.Get("/test", c.test)
	return r
}
