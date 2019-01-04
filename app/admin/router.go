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

type Message struct {
	AdminMessage string
}

func (s Service) test(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Message{"Congrats, you are admin!"})
}
