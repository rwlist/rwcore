package router

import (
	"github.com/rwlist/rwcore/cors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router *chi.Mux

type Deps struct {
	CORS cors.Middleware
}

func New(deps Deps) Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(deps.CORS)
	r.Use(a.Auth.Middleware)
	r.Use(a.DB.Middleware)

	r.Mount("/auth", a.Auth.Router())
	r.Mount("/admin", a.AdminService.Router())

	return r
}
