package model

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

type Roles map[string]struct{}

func (r Roles) HasRole(role string) bool {
	_, ok := r[role]
	return ok
}

func (r Roles) AddRole(role string) Roles {
	r[role] = struct{}{}
	return r
}

func (r Roles) RemoveRole(role string) Roles {
	delete(r, role)
	return r
}

func (r Roles) GetBSON() (interface{}, error) {
	arr := make([]string, 0)
	for k := range r {
		arr = append(arr, k)
	}
	return arr, nil
}

func (r *Roles) SetBSON(raw bson.Raw) error {
	initMap(r)
	var arr []string
	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}
	for _, v := range arr {
		r.AddRole(v)
	}
	return nil
}

func (r Roles) MarshalJSON() ([]byte, error) {
	arr := make([]string, 0)
	for k := range r {
		arr = append(arr, k)
	}
	return json.Marshal(arr)
}

func (r *Roles) UnmarshalJSON(b []byte) error {
	initMap(r)
	var arr []string
	err := json.Unmarshal(b, &arr)
	if err != nil {
		return err
	}
	for _, v := range arr {
		r.AddRole(v)
	}
	return nil
}

func initMap(i interface{}) {
	rv := reflect.ValueOf(i).Elem()
	t := rv.Type()
	rv.Set(reflect.MakeMap(t))
}
