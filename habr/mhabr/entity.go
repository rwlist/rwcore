package mhabr

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type ModelDaily struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Article   ArticleDaily  `bson:"article"`
	ArticleID string        `bson:"articleID"`
	Added     time.Time     `bson:"added" json:"added"`
}
