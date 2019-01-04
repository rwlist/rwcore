package main

import (
	"github.com/rwlist/rwcore/app/config"
	"log"

	"github.com/rwlist/rwcore/app"
	_ "github.com/rwlist/rwcore/app/resp"
	"github.com/rwlist/rwcore/modules/articles"
	"github.com/rwlist/rwcore/modules/habr"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conf, err := config.NewConfig("conf/config.json")
	if err != nil {
		log.Fatal(err)
	}

	app := app.CreateApp()
	app.AddModule(&articles.Module{})
	app.AddModule(&habr.Module{})
	app.Start()
}
