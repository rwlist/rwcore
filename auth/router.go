package auth

import (
	"github.com/go-chi/chi"
)

type Router struct { *chi.Mux }

func NewRouter(c *Controller) Router {
	r := chi.NewRouter()
	r.Get("/status", c.status)
	r.Post("/login", c.login)
	r.Post("/signup", c.signup)
	return Router{r}
}
