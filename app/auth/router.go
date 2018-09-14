package auth

import (
	"net/http"

	"github.com/rwlist/rwcore/app/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TokenResponse struct {
	User  interface{}
	Token string
}

func (a *Auth) Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/status", a.status)
	r.Post("/login", a.login)
	r.Post("/signup", a.signup)
	return r
}

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	claims, err := a.handleLogin(r)
	a.claimsResponse(w, r, claims, err)
}

func (a *Auth) signup(w http.ResponseWriter, r *http.Request) {
	claims, err := a.handleSignUp(r)
	a.claimsResponse(w, r, claims, err)
}

func (a *Auth) status(w http.ResponseWriter, r *http.Request) {
	claims, err := a.GetClaims(r)
	if err != nil {
		render.Render(w, r, utils.ErrUnathorized.With(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, TokenResponse{
		User:  claims.User,
		Token: getToken(r),
	})
}
