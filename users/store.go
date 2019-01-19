package users

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/cxt"
	"net/http"
)

const (
	CollName = "users"
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

func (u Store) Init() error {
	err1 := u.EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	})
	if err1 != nil {
		return err1
	}
	err2 := u.EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
	if err2 != nil {
		return err2
	}
	return nil
}

func (u Store) FindByUsername(username string) (user User, err error) {
	err = u.Find(bson.M{"username": username}).One(&user)
	return
}

func (u Store) InsertOne(user User) (err error) {
	err = user.Validate()
	if err != nil {
		return
	}
	err = u.Insert(user)
	return
}
