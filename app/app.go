package app

import (
	"github.com/rwlist/rwcore/conf"
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

func NewApp(conf conf.Server, db *db.Provider, auth *auth.Auth, adminService *admin.Service) (*App, error) {
	app := &App{
		DB:           db,
		Auth:         auth,
		AdminService: adminService,
		bindAddr:     conf.BindAddr,
	}
	app.Router = app.createRouter()

	err := app.DB.Init()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Start() error {
	log.Println("Server started")
	return http.ListenAndServe(a.bindAddr, a.Router)
}
