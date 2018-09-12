package db

import (
	"github.com/asaskevich/govalidator"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// TODO: consider using custom username validator

type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty" valid:"-"`
	Username       string        `bson:"username" valid:"utfletternum,runelength(3|30)"`
	HashedPassword []byte        `json:"-" bson:"password" valid:"-"`
	Email          string        `bson:"email" valid:"email"`
	FirstName      string        `bson:"firstName" valid:"utfletternum,runelength(3|30)"`
	SecondName     string        `bson:"secondName" valid:"utfletternum,runelength(3|30)"`
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

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

func (u UserStore) FindByUsername(username string) (user *User, err error) {
	err = u.c.Find(bson.M{"username": username}).One(&user)
	return
}

func (u UserStore) InsertOne(user *User) (err error) {
	err = user.Validate()
	if err != nil {
		return
	}
	err = u.c.Insert(user)
	return
}
