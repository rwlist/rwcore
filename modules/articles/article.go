package articles

import (
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/model"
)

type ArticleUpdate struct {
	ArticleID bson.ObjectId `json:"id"`
	Article   model.Article `json:"article"`
}

type ArticlesActionResp struct {
	AddedArticles []model.Article `json:"addedArticles"`
	Errors        []error         `json:"errors"`
}
