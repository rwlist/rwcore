package db

func (p *Provider) Users() UserStore {
	return UserStore{p.c("users")}
}
