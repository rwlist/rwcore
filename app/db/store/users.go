package store

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/model"
)

type Users struct {
	*mgo.Collection
}

func (u Users) Init() error {
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

func (u Users) FindByUsername(username string) (user *model.User, err error) {
	err = u.Find(bson.M{"username": username}).One(&user)
	return
}

func (u Users) InsertOne(user *model.User) (err error) {
	err = user.Validate()
	if err != nil {
		return
	}
	err = u.Insert(user)
	return
}
