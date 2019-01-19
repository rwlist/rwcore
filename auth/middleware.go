package auth

import (
	"context"
	"errors"
	"github.com/rwlist/rwcore/cxt"
	"github.com/rwlist/rwcore/users"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/resp"
)

const (
	BearerPrefix = "Bearer "
)

var (
	ErrTokenNotFound = errors.New("JWT token not found")
)

type Middleware struct {
	jwt *JWT
}

func NewMiddleware(jwt *JWT) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}

func (m *Middleware) ParseToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, BearerPrefix) {
		return "", ErrTokenNotFound
	}
	token := strings.TrimPrefix(auth, BearerPrefix)
	return token, nil
}

func (m *Middleware) ParseClaims(r *http.Request) *Claims {
	token, err := m.ParseToken(r)
	if err != nil {
		// ignore error, happens every first visit
		return nil
	}

	claims := Claims{}
	err = m.jwt.ParseAndVerify(token, &claims)
	if err != nil {
		log.Println("JWT token invalid:", err)
		return nil
	}

	return &claims
}

func (m *Middleware) UpdateContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims := m.ParseClaims(r)
		if claims != nil {
			ctx = context.WithValue(ctx, cxt.ClaimsKey, &claims)
			ctx = context.WithValue(ctx, cxt.UserKey, &claims.User)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func (m *Middleware) Authorized(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Context().Value(cxt.UserKey).(*users.User)
		if !ok {
			render.Render(w, r, resp.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (m *Middleware) HasRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(cxt.UserKey).(*users.User)
			if !ok {
				render.Render(w, r, resp.ErrUnauthorized)
				return
			}

			// Check if user has role or admin
			if !user.Roles.HasRole(role) && !user.Roles.HasRole(users.AdminRole) {
				render.Render(w, r, resp.ErrForbidden)
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
