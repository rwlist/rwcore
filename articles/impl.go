package articles

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/rwlist/rwcore/app/db"
	"github.com/rwlist/rwcore/app/model"
	"github.com/rwlist/rwcore/utils"
)

type impl struct{}

func (i impl) getAll(r *http.Request) ([]model.Article, error) {
	db := db.From(r)
	return db.Articles().GetAll()
}

func (i impl) addURL(url string, r *http.Request) (*ArticlesActionResp, error) {
	db := db.From(r)

	resp := &ArticlesActionResp{}

	title, err := utils.ParseTitleByURL(url)
	if err != nil {
		resp.Errors = append(resp.Errors, err)
	}

	if title == "" {
		title = "Unknown HTML Title"
	}

	article := model.NewArticle(url, title, map[string]interface{}{
		"added": "url",
	})
	db.Articles().InsertOne(&article)

	resp.AddedArticles = append(resp.AddedArticles, article)

	return resp, nil
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
		return nil, err
	}

	article.Status.Rating += delta

	err = db.Articles().UpdateOne(&article)
	return asUpdate(article), err
}

func (i impl) removeTag(r *http.Request, article model.Article) (*ArticleUpdate, error) {
	// TODO: queue
	db := db.From(r)
	now := time.Now()

	tag := r.URL.Query().Get("tag")
	err := article.RemoveTag(tag)
	if err != nil {
		return nil, err
	}
	article.Modified = &now

	err = db.Articles().UpdateOne(&article)
	return asUpdate(article), err
}

func (i impl) addTag(r *http.Request, article model.Article) (*ArticleUpdate, error) {
	// TODO: queue
	db := db.From(r)
	now := time.Now()

	tag := r.URL.Query().Get("tag")
	err := article.AddTag(tag)
	if err != nil {
		return nil, err
	}
	article.Modified = &now

	err = db.Articles().UpdateOne(&article)
	return asUpdate(article), err
}

func asUpdate(article model.Article) *ArticleUpdate {
	return &ArticleUpdate{
		ArticleID: article.ID,
		Article:   article,
	}
}
