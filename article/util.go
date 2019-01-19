package article

import (
	"github.com/rwlist/rwcore/cxt"
	"net/http"
)

func asUpdate(article Article) *ArticleUpdate {
	return &ArticleUpdate{
		ArticleID: article.ID,
		Article:   article,
	}
}

func R(r *http.Request) Article {
	return r.Context().Value(cxt.ArticleKey).(Article)
}
