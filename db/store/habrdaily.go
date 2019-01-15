package store

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/model"
)

type HabrDaily struct {
	*mgo.Collection
}

func (s HabrDaily) CountByArticleID(ID string) (int, error) {
	return s.Find(bson.M{"articleID": ID}).Count()
}

func (s HabrDaily) InsertOne(article model.HabrDailyArticle) (model.HabrDailyArticle, error) {
	err := s.Insert(&article)
	return article, err
}
