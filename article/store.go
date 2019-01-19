package article

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/cxt"
	"net/http"
)

const (
	CollName = "articles"
)

func S(db *mgo.Database) Store {
	return Store{db.C(CollName)}
}

func DB(r *http.Request) Store {
	return S(cxt.DB(r.Context()))
}

type Store struct {
	*mgo.Collection
}

func (s Store) Init() error {
	iter := s.Find(bson.M{}).Iter()
	var article Article
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

func (s Store) GetAll() ([]Article, error) {
	var articles []Article
	err := s.Find(bson.M{}).All(&articles)
	return articles, err
}

func (s Store) InsertOne(article *Article) (err error) {
	err = s.Insert(article)
	return
}

func (s Store) InsertMany(articles []Article) (err error) {
	iSlice := make([]interface{}, len(articles))
	for k, v := range articles {
		iSlice[k] = v
	}
	err = s.Insert(iSlice...)
	return
}

func (s Store) UpdateOne(article *Article) (err error) {
	return s.UpdateId(article.ID, article)
}
