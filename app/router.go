package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(a.Auth.Middleware)
	r.Use(a.DB.Middleware)

	r.Mount("/auth", a.Auth.Router())

	return r
}
