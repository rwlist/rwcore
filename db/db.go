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

func NewProvider(conf Config) (*Provider, error) {
	session, err := mgo.Dial(conf.Addr)
	if err != nil {
		return nil, err
	}
	return &Provider{session, conf.DbName}, nil
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
