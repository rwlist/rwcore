package habr

import (
	"github.com/globalsign/mgo"
	"github.com/rwlist/rwcore/cxt"
	"net/http"
)

const (
	CollName = "habrDaily"
)

func DB(r *http.Request) DailyStore {
	return DailyStore{
		cxt.DB(r.Context()).C(CollName),
	}
}

type DailyStore struct {
	*mgo.Collection
}
