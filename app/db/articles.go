package db

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/model"
)

type ArticleStore struct {
	c *mgo.Collection
}

func (s ArticleStore) EnsureIndexes() error {
	return nil
}

func (s ArticleStore) GetAll() ([]model.Article, error) {
	var articles []model.Article
	err := s.c.Find(bson.M{}).All(&articles)
	return articles, err
}

func (s ArticleStore) UpdateTags(id bson.ObjectId, tags map[string]string) error {
	return s.c.Update(
		bson.M{
			"_id": id,
		},
		bson.M{
			"tags": tags,
		},
	)
}

func (s ArticleStore) InsertOne(article *model.Article) error {
	return s.c.Insert(article)
}

func (s ArticleStore) Size() (int, error) {
	return s.c.Count()
}
