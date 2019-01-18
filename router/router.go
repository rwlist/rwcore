package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/wire"
	"github.com/rwlist/rwcore/admin"
	"github.com/rwlist/rwcore/auth"
	"github.com/rwlist/rwcore/cors"
	"github.com/rwlist/rwcore/mod"
)

type Router struct{ *chi.Mux }

type Mid struct {
	CORS cors.Middleware
	Auth *auth.Middleware
	DB   *mod.Middleware
}

type Routes struct {
	Auth  auth.Router
	Admin admin.Router
}

func New(mid *Mid, rt *Routes) Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(mid.CORS)
	r.Use(mid.Auth.UpdateContext)
	r.Use(mid.DB.UpdateContext)

	r.Mount("/auth", rt.Auth)
	r.Mount("/admin", rt.Admin)

	return Router{r}
}

var All = wire.NewSet(
	New,
	Mid{},
	Routes{},
)
