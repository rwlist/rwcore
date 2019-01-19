package article

import (
	"context"
	"errors"
	"github.com/rwlist/rwcore/cxt"
	"github.com/rwlist/rwcore/resp"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) UpdateContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		articles := DB(r)

		id := chi.URLParam(r, "id")
		if !bson.IsObjectIdHex(id) {
			render.Render(w, r, resp.ErrBadRequest.With(errors.New("invalid bson object id hex")))
			return
		}
		bid := bson.ObjectIdHex(id)

		var article Article
		err := articles.FindId(bid).One(&article)
		if err != nil {
			render.Render(w, r, resp.ErrNotFound.With(err))
			return
		}

		ctx = context.WithValue(ctx, cxt.ArticleKey, article)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
