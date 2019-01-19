//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rwlist/rwcore/admin"
	"github.com/rwlist/rwcore/article"
	"github.com/rwlist/rwcore/auth"
	"github.com/rwlist/rwcore/conf"
	"github.com/rwlist/rwcore/cors"
	"github.com/rwlist/rwcore/habr"
	"github.com/rwlist/rwcore/mod"
	"github.com/rwlist/rwcore/mod/dbinit"
	"github.com/rwlist/rwcore/router"
	"github.com/rwlist/rwcore/srv"
)

func Initialize(filepath string) (App, func(), error) {
	wire.Build(
		conf.All,
		admin.All,
		mod.All,
		dbinit.All,
		cors.All,
		router.All,
		srv.All,
		auth.All,
		article.All,
		habr.All,
		App{},
	)
	return App{}, nil, nil
}
