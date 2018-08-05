package mongodb

import "github.com/globalsign/mgo"

var MaxPool int
var PATH string
var DBNAME string

var BaseSession *mgo.Session

func CheckAndInitServiceConnection() {
	if BaseSession == nil {
		var err error
		BaseSession, err = mgo.Dial(PATH)
		if err != nil {
			panic(err)
		}
	}
	if service.baseSession == nil {
		service.URL = PATH
		err := service.New()
		if err != nil {
			panic(err)
		}
	}
}
