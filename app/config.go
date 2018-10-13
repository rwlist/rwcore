package app

import (
	"github.com/rwlist/rwcore/app/auth"
	"github.com/rwlist/rwcore/app/db"
)

type RootConfig struct {
	Server Server
	Mongo  db.Config
	Auth   auth.Config
}

type Server struct {
	BindAddr string
}
