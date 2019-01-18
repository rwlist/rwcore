package store

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/model"
)

type Articles struct {
	*mgo.Collection
}

func (s Articles) Init() error {
	iter := s.Find(bson.M{}).Iter()
	var article model.Article
	for iter.Next(&article) {
		ok, err := article.BumpVersion()
		if err != nil {
			iter.Close()
			return err
		}
		if !ok {
			continue
		}
		err = s.UpdateOne(&article)
		if err != nil {
			return err
		}
	}
	return iter.Close()
}

func (s Articles) GetAll() ([]model.Article, error) {
	var articles []model.Article
	err := s.Find(bson.M{}).All(&articles)
	return articles, err
}

func (s Articles) InsertOne(article *model.Article) (err error) {
	err = s.Insert(article)
	return
}

func (s Articles) InsertMany(articles []model.Article) (err error) {
	iSlice := make([]interface{}, len(articles))
	for k, v := range articles {
		iSlice[k] = v
	}
	err = s.Insert(iSlice...)
	return
}

func (s Articles) UpdateOne(article *model.Article) (err error) {
	return s.UpdateId(article.ID, article)
}
