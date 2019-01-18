package main

import (
	"flag"
	"github.com/rwlist/rwcore/mod"
	"log"

	"github.com/rwlist/rwcore/srv"
)

type App struct {
	Server *srv.Server
	DbInit *mod.Init
}

func main() {
	log.SetFlags(log.Lshortfile)

	filepath := *flag.String("config", "conf/config.toml", "pass path to config file")
	flag.Parse()

	app, cleanup, err := Initialize(filepath)
	defer cleanup()
	if err != nil {
		log.Println("Initialization failed.", err)
		return
	}


	err = app.DbInit.Do()
	if err != nil {
		log.Println("DB init failed.", err)
		return
	}
	app.Server.Start()
}
