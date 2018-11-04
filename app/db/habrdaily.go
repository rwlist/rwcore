package db

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/model"
)

type HabrDailyStore struct {
	c *mgo.Collection
}

func (s HabrDailyStore) EnsureIndexes() error {
	return nil
}

func (s HabrDailyStore) CountByArticleID(ID string) (int, error) {
	return s.c.Find(bson.M{"articleID": ID}).Count()
}

func (s HabrDailyStore) InsertOne(article model.HabrDailyArticle) (model.HabrDailyArticle, error) {
	err := s.c.Insert(&article)
	return article, err
}

func (s HabrDailyStore) Size() (int, error) {
	return s.c.Count()
}
