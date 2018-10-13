package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func (a *App) createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(a.Auth.Middleware)
	r.Use(a.DB.Middleware)

	r.Mount("/auth", a.Auth.Router())
	r.Mount("/admin", a.AdminService.Router())

	return r
}
