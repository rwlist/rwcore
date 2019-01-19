package auth

import (
	"encoding/json"
	"github.com/rwlist/rwcore/users"
	"github.com/rwlist/rwcore/resp"
	"net/http"

	"github.com/go-chi/render"
)

type Controller struct{
	auth *Auth
	jwt *JWT
	middleware *Middleware
}

func NewController(auth *Auth, jwt *JWT, middleware *Middleware) *Controller {
	c := &Controller{
		auth: auth,
		jwt: jwt,
		middleware: middleware,
	}
	return c
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	db := users.DB(r)

	decoder := json.NewDecoder(r.Body)
	var form LoginForm
	err := decoder.Decode(&form)
	if err != nil {
		render.Render(w, r, resp.ErrBadRequest.With(err))
		return
	}

	user, err := c.auth.Login(db, form)
	if err != nil {
		render.Render(w, r, resp.ErrBadRequest.With(err))
		return
	}

	claims := c.auth.CreateClaims(user)

	token, err := c.jwt.CreateAndSign(claims)
	if err != nil {
		render.Render(w, r, resp.ErrInternal.With(err))
		return
	}

	resp := TokenResponse{
		User: user,
		Token: token,
	}

	render.Respond(w, r, resp)
}

func (c *Controller) signup(w http.ResponseWriter, r *http.Request) {
	db := users.DB(r)

	decoder := json.NewDecoder(r.Body)
	var form SignUpForm
	err := decoder.Decode(&form)
	if err != nil {
		render.Render(w, r, resp.ErrBadRequest.With(err))
		return
	}

	user, err := c.auth.SignUp(db, form)
	if err != nil {
		render.Render(w, r, resp.ErrBadRequest.With(err))
		return
	}

	claims := c.auth.CreateClaims(user)

	token, err := c.jwt.CreateAndSign(claims)
	if err != nil {
		render.Render(w, r, resp.ErrInternal.With(err))
		return
	}

	resp := TokenResponse{
		User: user,
		Token: token,
	}

	render.Respond(w, r, resp)
}

func (c *Controller) status(w http.ResponseWriter, r *http.Request) {
	claims := c.middleware.ParseClaims(r)

	if claims == nil {
		render.Render(w, r, resp.ErrUnauthorized)
		return
	}

	token, _ := c.middleware.ParseToken(r)

	render.Respond(w, r, TokenResponse{
		User:  claims.User,
		Token: token,
	})
}
