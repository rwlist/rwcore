package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/app/model"
	"github.com/rwlist/rwcore/app/resp"
)

type Claims struct {
	User *model.User
	jwt.StandardClaims
}

func (a *Auth) newClaims(user *model.User) (*Claims, error) {
	cur := time.Now().UTC()
	exp := cur.Add(a.jwtDuration)
	return &Claims{
		StandardClaims: jwt.StandardClaims{
			NotBefore: cur.Unix(),
			IssuedAt:  cur.Unix(),
			ExpiresAt: exp.Unix(),
		},
		User: user,
	}, nil
}

func (a *Auth) claimsResponse(w http.ResponseWriter, r *http.Request, claims *Claims, err error) {
	if err != nil {
		render.Render(w, r, resp.ErrBadRequest.With(err))
		return
	}
	token, err := a.createToken(claims)
	if err != nil {
		render.Render(w, r, resp.ErrInternal.With(err))
	}
	render.Respond(w, r, TokenResponse{claims.User, token})
}
