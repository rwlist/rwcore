package db

import (
	"github.com/globalsign/mgo"
)

type Config struct {
	Addr   string
	DbName string
}

type Provider struct {
	session *mgo.Session
	dbName  string
}

func New(conf Config) (*Provider, error) {
	session, err := mgo.Dial(conf.Addr)
	if err != nil {
		return nil, err
	}
	provider := &Provider{session, conf.DbName}
	err = provider.EnsureIndexes()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (p *Provider) EnsureIndexes() error {
	err := p.Users().EnsureIndexes()
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) Copy() *Provider {
	return &Provider{p.session.Copy(), p.dbName}
}

func (p *Provider) Close() {
	p.session.Close()
	p.session = nil
}

func (p *Provider) db() *mgo.Database {
	return p.session.DB(p.dbName)
}

func (p *Provider) c(name string) *mgo.Collection {
	return p.db().C(name)
}
