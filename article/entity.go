package article

import (
	"github.com/globalsign/mgo/bson"
	"github.com/google/wire"
)

type ArticleUpdate struct {
	ArticleID bson.ObjectId `json:"id"`
	Article   Article       `json:"article"`
}

type ArticlesActionResp struct {
	AddedArticles []Article `json:"addedArticles"`
	Errors        []error   `json:"errors"`
}

var All = wire.NewSet( // TODO
	NewController,
	NewMiddleware,
	NewRouter,
	NewService,
)
