package model

import (
	"github.com/globalsign/mgo/bson"
)

type HabrArticlePreview struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	
}
