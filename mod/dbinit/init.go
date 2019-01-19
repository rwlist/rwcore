package dbinit

import (
	"github.com/google/wire"
	"github.com/rwlist/rwcore/article"
	"github.com/rwlist/rwcore/mod"
	"github.com/rwlist/rwcore/users"
)

type Once struct {
	Provider *mod.Provider
}

func NewOnce(provider *mod.Provider) *Once {
	return &Once{
		Provider: provider,
	}
}

func (o *Once) Do() error {
	_, db, cleanup := o.Provider.Copy()
	defer cleanup()

	var err error

	err = users.S(db).Init()
	if err != nil {
		return err
	}

	err = article.S(db).Init()
	if err != nil {
		return err
	}

	return nil
}

var All = wire.NewSet(NewOnce)
