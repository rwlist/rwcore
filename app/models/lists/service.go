package lists

import (
	"github.com/globalsign/mgo"
)

var DBNAME string = "rwlists"
var DBPATH string

type Service struct {
	baseSession *mgo.Session
}

var service Service

func (s *Service) Init() error {
	session, err := mgo.Dial(DBPATH)
	if err != nil {
		return err
	}
	s.baseSession = session
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
