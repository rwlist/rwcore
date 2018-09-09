package lists

import (
	"github.com/globalsign/mgo"
	"github.com/rwlist/rwcore/app/models/mongodb"
)

var DBNAME string = "rwlists"
var DBPATH string

type Service struct {
	baseSession *mgo.Session
}

var service Service

func (s *Service) Init() error {
	s.baseSession = mongodb.BaseSession
	return nil
}

func (s *Service) DB() Database {
	session := s.baseSession.Clone()
	return Database{session.DB(DBNAME), session}
}

func InitService() {
	if service.baseSession == nil {
		err := service.Init()
		if err != nil {
			panic(err)
		}
	}
}
