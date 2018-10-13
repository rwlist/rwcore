package admin

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/app/auth"
)

func (s Service) Router() http.Handler {
	r := chi.NewRouter()
	r.Use(auth.HasRole("admin"))
	r.Get("/test", s.test)
	return r
}

type AdminMessage struct {
	AdminMessage string
}

func (s Service) test(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, AdminMessage{"Congrats, you are admin!"})
}
