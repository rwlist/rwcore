package db

func (p *Provider) Users() UserStore {
	return UserStore{p.c("users")}
}

func (p *Provider) Articles() ArticleStore {
	return ArticleStore{p.c("articles")}
}

func (p *Provider) HabrDaily() HabrDailyStore {
	return HabrDailyStore{p.c("habrDaily")}
}
