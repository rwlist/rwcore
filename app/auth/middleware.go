package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/app/model"
	"github.com/rwlist/rwcore/app/utils"
)

type key int

const (
	UserKey key = iota
	ClaimsKey
)

func (a *Auth) Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		claims, err := a.GetClaims(r)
		if claims != nil {
			ctx = context.WithValue(ctx, ClaimsKey, claims)
			ctx = context.WithValue(ctx, UserKey, claims.User)
		}
		if err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func Authorized(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Context().Value(UserKey).(*model.User)
		if !ok {
			render.Render(w, r, utils.ErrUnathorized)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func HasRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(UserKey).(*model.User)
			if !ok {
				render.Render(w, r, utils.ErrUnathorized)
				return
			}
			// Check if user has role or admin
			if !user.Roles.HasRole(role) && !user.Roles.HasRole(AdminRole) {
				render.Render(w, r, utils.ErrForbidden)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
