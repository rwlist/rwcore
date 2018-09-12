package auth

import (
	"net/http"

	"github.com/go-chi/chi"
	. "github.com/rwlist/rwcore/app/utils"
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
	a.processUser(w, user, err)
}

func (a *Auth) signup(w http.ResponseWriter, r *http.Request) {
	user, err := a.provider.CreateUser(r)
	a.processUser(w, user, err)
}

func (a *Auth) status(w http.ResponseWriter, r *http.Request) {
	user, err := a.GetUser(r)
	if err != nil {
		WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}
	WriteJSON(w, http.StatusOK, TokenResponse{
		User:  user,
		Token: getToken(r),
	})
}
