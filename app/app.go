package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rwlist/rwcore/app/admin"
	"github.com/rwlist/rwcore/app/auth"
	"github.com/rwlist/rwcore/app/db"
)

type App struct {
	DB           *db.Provider
	Auth         *auth.Auth
	AdminService *admin.Service

	Router   *chi.Mux
	bindAddr string
}

func CreateApp(conf RootConfig) *App {
	var err error

	app := &App{}
	app.DB, err = db.New(conf.Mongo)
	if err != nil {
		log.Fatal(err)
	}
	app.Auth, err = auth.New(conf.Auth)
	if err != nil {
		log.Fatal(err)
	}
	app.AdminService = &admin.Service{}

	app.Router = app.createRouter()
	app.bindAddr = conf.Server.BindAddr

	return app
}

func (app *App) Start() error {
	return http.ListenAndServe(app.bindAddr, app.Router)
}
