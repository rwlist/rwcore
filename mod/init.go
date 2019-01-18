package mod

import "github.com/rwlist/rwcore/users"

type Init struct{
	Provider *Provider
}

func NewInit(provider *Provider) *Init {
	return &Init{
		Provider: provider,
	}
}

func (i *Init) Do() error {
	_, db, cleanup := i.Provider.Copy()
	defer cleanup()

	var err error

	err = users.Store{ Collection: db.C(users.CollName) }.Init()
	if err != nil {
		return err
	}

	return nil
}