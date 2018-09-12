package db

import (
	"context"
	"net/http"
)

const (
	DBKey = "DB"
)

func (p *Provider) Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		db := p.Copy()
		defer db.Close()

		ctx := r.Context()
		ctx = context.WithValue(ctx, DBKey, db)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
