package db

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/model"
)

type UserStore struct {
	c *mgo.Collection
}

func (u UserStore) EnsureIndexes() error {
	err1 := u.c.EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	})
	if err1 != nil {
		return err1
	}
	err2 := u.c.EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
	if err2 != nil {
		return err2
	}
	return nil
}

func (u UserStore) FindByUsername(username string) (user *model.User, err error) {
	err = u.c.Find(bson.M{"username": username}).One(&user)
	return
}

func (u UserStore) InsertOne(user *model.User) (err error) {
	err = user.Validate()
	if err != nil {
		return
	}
	err = u.c.Insert(user)
	return
}

func (u UserStore) Size() (int, error) {
	return u.c.Count()
}
