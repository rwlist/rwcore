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

// UpdateTags -
//
// unused, todo: remove
func (s ArticleStore) UpdateTags(id bson.ObjectId, tags map[string]string) error {
	return s.c.Update(
		bson.M{
			"_id": id,
		},
		bson.M{
			"$set": bson.M{
				"tags": tags,
			},
		},
	)
}

func (s ArticleStore) InsertOne(article *model.Article) (err error) {
	err = article.BeforeInsert()
	if err != nil {
		return err
	}
	err = s.c.Insert(article)
	return
}

func (s ArticleStore) InsertMany(articles []model.Article) (err error) {
	for k := range articles {
		err = articles[k].BeforeInsert()
		if err != nil {
			return err
		}
	}
	err = s.c.Insert(articles)
	return
}

func (s ArticleStore) UpdateOne(article *model.Article) (err error) {
	return s.c.UpdateId(article.ID, article)
}

func (s ArticleStore) Size() (int, error) {
	return s.c.Count()
}
