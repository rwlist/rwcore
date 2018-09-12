package auth

import (
	"context"
	"net/http"
)

const (
	UserKey   = "User"
	ClaimsKey = "Claims"
)

func (a *Auth) Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		claims, _ := a.GetClaims(r)
		user, _ := a.GetUser(r)

		ctx := r.Context()
		if claims != nil {
			ctx = context.WithValue(ctx, ClaimsKey, claims)
		}
		if user != nil {
			ctx = context.WithValue(ctx, UserKey, user)
		}
		
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
