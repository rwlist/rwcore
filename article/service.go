package article

import (
	"errors"
	"time"

	"github.com/rwlist/rwcore/utils"
)

// TODO: queue/mutex
type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) AddURL(db Store, url string) (*ArticlesActionResp, error) {
	resp := &ArticlesActionResp{}

	title, err := utils.ParseTitleByURL(url)
	if err != nil {
		resp.Errors = append(resp.Errors, err)
	}

	if title == "" {
		title = "Unknown HTML Title"
	}

	article := NewArticle(url, title, map[string]interface{}{
		"added": "url",
	})

	err = db.InsertOne(&article)
	if err != nil {
		return nil, err
	}

	resp.AddedArticles = append(resp.AddedArticles, article)

	return resp, nil
}

func (s Service) OnClick(db Store, article Article) (*ArticleUpdate, error) {
	now := time.Now()

	article.Status.Clicks++
	article.Status.LastClick = &now
	if article.Status.ReadStatus == "unopened" {
		article.Status.ReadStatus = "viewed"
		article.Status.ReadStatusChange = &now
	}

	err := db.UpdateOne(&article)
	return asUpdate(article), err
}

func (s Service) SetReadStatus(db Store, article Article, newStatus string) (*ArticleUpdate, error) {
	now := time.Now()

	if !ValidArticleReadStatus(newStatus) {
		return nil, errors.New("invalid read status")
	}

	article.Status.ReadStatus = newStatus
	article.Status.ReadStatusChange = &now

	err := db.UpdateOne(&article)
	return asUpdate(article), err
}

func (s Service) ChangeRating(db Store, article Article, delta int) (*ArticleUpdate, error) {
	article.Status.Rating += delta

	err := db.UpdateOne(&article)
	return asUpdate(article), err
}

func (s Service) RemoveTag(db Store, article Article, tag string) (*ArticleUpdate, error) {
	now := time.Now()

	err := article.RemoveTag(tag)
	if err != nil {
		return nil, err
	}
	article.Modified = &now

	err = db.UpdateOne(&article)
	return asUpdate(article), err
}

func (s Service) AddTag(db Store, article Article, tag string) (*ArticleUpdate, error) {
	now := time.Now()

	err := article.AddTag(tag)
	if err != nil {
		return nil, err
	}
	article.Modified = &now

	err = db.UpdateOne(&article)
	return asUpdate(article), err
}
