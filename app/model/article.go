package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Article struct {
	ID    bson.ObjectId     `bson:"_id,omitempty" json:"id"`
	URL   string            `bson:"url" json:"url"`
	Added time.Time         `bson:"added" json:"added"`
	Tags  map[string]string `bson:"tags" json:"tags"`
}
