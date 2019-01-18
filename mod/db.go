package mod

import (
	"github.com/globalsign/mgo"
	"github.com/google/wire"
)

type Config struct {
	Addr   string
	DbName string
}

type Provider struct {
	session *mgo.Session
	dbName  string
}

func NewProvider(conf Config) (*Provider, func(), error) {
	session, err := mgo.Dial(conf.Addr)
	if err != nil {
		return nil, nil, err
	}

	return &Provider{
		session: session,
		dbName: conf.DbName,
	}, session.Close, nil
}

func (p *Provider) Copy() (*mgo.Session, *mgo.Database, func()) {
	session := p.session.Copy()
	db := session.DB(p.dbName)
	cleanup := session.Close
	return session, db, cleanup
}

var All = wire.NewSet(
	NewProvider,
	NewInit,
	NewMiddleware,
)