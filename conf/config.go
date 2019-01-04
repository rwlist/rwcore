package conf

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/rwlist/rwcore/app/auth"
	"github.com/rwlist/rwcore/app/db"
)

type Config struct {
	Server Server
	Mongo  db.Config
	Auth   auth.Config
}

type Server struct {
	BindAddr string
}

func New(filepath string) (Config, error) {
	conf := config.NewConfig()
	err := conf.Load(file.NewSource(
		file.WithPath(filepath),
	))
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = conf.Scan(&c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
