package articles

import (
	"net/http"

	"github.com/rwlist/rwcore/app/db"
	"github.com/rwlist/rwcore/app/model"
)

type impl struct{}

func (i impl) getAll(r *http.Request) ([]model.Article, error) {
	db := db.From(r)
	return db.Articles().GetAll()
}

// func addOneArticle(r *http.Request, article model.Article) (model.Article, error) {
// 	db := db.From(r)
// 	err := db.Articles().InsertOne(&article)
// 	return article, err
// }

// func addManyArticles(r *http.Request, articles []model.Article) ([]model.Article, error) {
// 	db := db.From(r)
// 	err := db.Articles().InsertMany(articles)
// 	return articles, err
// }

// func patchArticle(r *http.Request, article model.Article) (model.Article, error) {
// 	db := db.From(r)
// 	err := db.Articles().UpdateOne(&article)
// 	return article, err
// }
