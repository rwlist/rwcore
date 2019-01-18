package mod

import (
	"context"
	"github.com/rwlist/rwcore/cxt"
	"net/http"
)


type Middleware struct{
	provider *Provider
}

func NewMiddleware(provider *Provider) *Middleware {
	return &Middleware{
		provider: provider,
	}
}

func (m *Middleware) UpdateContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, db, cleanup := m.provider.Copy()
		defer cleanup()

		ctx := r.Context()
		ctx = context.WithValue(ctx, cxt.MgoKey ,session)
		ctx = context.WithValue(ctx, cxt.DbKey, db)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}