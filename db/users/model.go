package users

import (
	"encoding/json"
	"reflect"

	"github.com/asaskevich/govalidator"
	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty" valid:"-"`
	Username       string        `bson:"username" valid:"utfletternum,runelength(3|30)"`
	HashedPassword []byte        `json:"-" bson:"password" valid:"-"`
	Email          string        `bson:"email" valid:"email"`
	FirstName      string        `bson:"firstName" valid:"utfletternum,runelength(3|30)"`
	SecondName     string        `bson:"secondName" valid:"utfletternum,runelength(3|30)"`
	Roles          Roles         `bson:"roles" valid:"-"`
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}
