package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Article struct {
	ID    bson.ObjectId          `bson:"_id,omitempty" json:"id"`
	URL   string                 `bson:"url" json:"url"`
	Added time.Time              `bson:"added" json:"added"`
	Tags  map[string]interface{} `bson:"tags" json:"tags"`
}

func (a *Article) BeforeInsert() error {
	a.ID = bson.NewObjectId()
	a.Added = time.Now()
	return nil
}
