package conf

import (
	"github.com/google/wire"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/rwlist/rwcore/auth"
	"github.com/rwlist/rwcore/mod"
	"github.com/rwlist/rwcore/srv"
)

type Config struct {
	Server srv.Config
	Mongo  mod.Config
	JWT    auth.JWTConfig
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

func ProvideSrv(c Config) srv.Config {
	return c.Server
}

func ProvideMongo(c Config) mod.Config {
	return c.Mongo
}

func ProvideJWT(c Config) auth.JWTConfig {
	return c.JWT
}

var All = wire.NewSet(
	New,
	ProvideSrv,
	ProvideMongo,
	ProvideJWT,
)
