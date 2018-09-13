package auth

import (
	"net/http"

	"github.com/rwlist/rwcore/app/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (a *Auth) Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/status", a.status)
	r.Post("/login", a.login)
	r.Post("/signup", a.signup)
	return r
}

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	user, err := a.provider.FindUser(r)
	a.processUser(w, r, user, err)
}

func (a *Auth) signup(w http.ResponseWriter, r *http.Request) {
	user, err := a.provider.CreateUser(r)
	a.processUser(w, r, user, err)
}

func (a *Auth) status(w http.ResponseWriter, r *http.Request) {
	user, err := a.GetUser(r)
	if err != nil {
		render.Render(w, r, utils.ErrUnathorized.With(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, TokenResponse{
		User:  user,
		Token: getToken(r),
	})
}
