package habr

import (
	"github.com/rwlist/rwcore/article"
	"github.com/rwlist/rwcore/habr/client"
	"github.com/rwlist/rwcore/habr/mhabr"
	"github.com/rwlist/rwcore/mod"
)

type Service struct {
	reader *client.ReaderDailyTop

	habrDB    mhabr.StoreDaily
	articleDB article.Store
}

func NewService(reader *client.ReaderDailyTop, prov *mod.Provider) (*Service, func()) {
	_, db, cleanup := prov.Copy()
	return &Service{
		reader:    reader,
		habrDB:    mhabr.S(db),
		articleDB: article.S(db),
	}, cleanup
}
