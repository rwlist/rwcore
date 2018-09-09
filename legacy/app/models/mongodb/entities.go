package mongodb

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Username     string        `bson:"username"`
	PasswordHash []byte        `bson:"passwordHash"`
}
