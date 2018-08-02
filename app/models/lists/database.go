package lists

import (
	"github.com/globalsign/mgo"
)

type Database struct {
	*mgo.Database
	session *mgo.Session
}

func (d Database) Close() {
	d.session.Close()
}
