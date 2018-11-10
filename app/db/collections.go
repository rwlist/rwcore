package db

import (
	"github.com/rwlist/rwcore/app/db/store"
)

func (p *Provider) Users() store.Users {
	return store.Users{p.c("users")}
}

func (p *Provider) Articles() store.Articles {
	return store.Articles{p.c("articles")}
}

func (p *Provider) HabrDaily() store.HabrDaily {
	return store.HabrDaily{p.c("habrDaily")}
}

func (p *Provider) AllCollections() []interface{} {
	return []interface{}{
		p.Users(),
		p.Articles(),
		p.HabrDaily(),
	}
}
