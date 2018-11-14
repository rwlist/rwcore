package articles

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/rwlist/rwcore/app/db"
	"github.com/rwlist/rwcore/app/model"
)

type impl struct{}

func (i impl) getAll(r *http.Request) ([]model.Article, error) {
	db := db.From(r)
	return db.Articles().GetAll()
}

func (i impl) onClick(r *http.Request, article model.Article) (*ArticleUpdate, error) {
	// TODO: queue
	db := db.From(r)
	now := time.Now()

	article.Status.Clicks++
	article.Status.LastClick = &now
	if article.Status.ReadStatus == "unopened" {
		article.Status.ReadStatus = "viewed"
		article.Status.ReadStatusChange = &now
	}

	err := db.Articles().UpdateOne(&article)
	return asUpdate(article), err
}

func (i impl) setReadStatus(r *http.Request, article model.Article) (*ArticleUpdate, error) {
	// TODO: queue
	db := db.From(r)
	now := time.Now()

	newStatus := r.URL.Query().Get("newStatus")
	if !model.ValidArticleReadStatus(newStatus) {
		return nil, errors.New("invalid read status")
	}

	article.Status.ReadStatus = newStatus
	article.Status.ReadStatusChange = &now

	err := db.Articles().UpdateOne(&article)
	return asUpdate(article), err
}

func (i impl) changeRating(r *http.Request, article model.Article) (*ArticleUpdate, error) {
	// TODO: queue
	db := db.From(r)

	delta, err := strconv.Atoi(r.URL.Query().Get("delta"))
	if err != nil {
		return nil, errors.New("invalid delta")
	}

	article.Status.Rating += delta

	err = db.Articles().UpdateOne(&article)
	return asUpdate(article), err
}

func asUpdate(article model.Article) *ArticleUpdate {
	return &ArticleUpdate{
		ArticleID: article.ID,
		Article:   article,
	}
}
