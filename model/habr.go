package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/habr/models"
)

type HabrDailyArticle struct {
	ID        bson.ObjectId       `bson:"_id,omitempty"`
	Article   models.ArticleDaily `bson:"article"`
	ArticleID string              `bson:"articleID"`
	Added     time.Time           `bson:"added" json:"added"`
}
