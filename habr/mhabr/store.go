package mhabr

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	CollNameDaily = "habrDaily"
)

func S(db *mgo.Database) StoreDaily {
	return StoreDaily{db.C(CollNameDaily)}
}

type StoreDaily struct {
	*mgo.Collection
}

func (s StoreDaily) CountByArticleID(ID string) (int, error) {
	return s.Find(bson.M{"articleID": ID}).Count()
}

func (s StoreDaily) InsertOne(article ModelDaily) (ModelDaily, error) {
	err := s.Insert(&article)
	return article, err
}
