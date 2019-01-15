//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rwlist/rwcore/admin"
	"github.com/rwlist/rwcore/conf"
	"github.com/rwlist/rwcore/cors"
	"github.com/rwlist/rwcore/router"
	"github.com/rwlist/rwcore/srv"
)

func Initialize(filepath string) (App, error) {
	wire.Build(
		// config
		conf.All,

		// services
		admin.All,

		// middlewares
		cors.NewMiddleware,

		// router
		router.Deps,
		router.New,

		srv.Deps,
		srv.New,

		App{},
	)
	return App{}, nil
}
