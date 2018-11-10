package articles

import (
	"context"
	"errors"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/app/db"
	"github.com/rwlist/rwcore/app/model"
	"github.com/rwlist/rwcore/app/utils"
)

type ctxkey string

var (
	articleKey ctxkey = "article"
)

func fetchArticle(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		articles := db.From(r).Articles()

		id := chi.URLParam(r, "id")
		if !bson.IsObjectIdHex(id) {
			render.Render(w, r, utils.ErrBadRequest.With(errors.New("invalid bson object id hex")))
			return
		}
		bid := bson.ObjectIdHex(id)

		var article model.Article
		err := articles.FindId(bid).One(&article)
		if err != nil {
			render.Render(w, r, utils.ErrNotFound.With(err))
			return
		}

		ctx = context.WithValue(ctx, articleKey, article)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
