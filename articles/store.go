package articles

import (
	"github.com/globalsign/mgo"
	"github.com/rwlist/rwcore/cxt"
	"net/http"
)

const (
	CollName = "articles"
)

func DB(r *http.Request) Store {
	return Store{
		cxt.DB(r.Context()).C(CollName),
	}
}

type Store struct {
	*mgo.Collection
}
