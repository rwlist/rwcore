package main

import (
	"log"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/rwlist/rwcore/app"
)

func main() {
	conf := config.NewConfig()
	err := conf.Load(file.NewSource(
		file.WithPath("conf/config.json"),
	))
	if err != nil {
		log.Fatal(err)
	}

	var root app.RootConfig
	err = conf.Scan(&root)
	if err != nil {
		log.Fatal(err)
	}

	app := app.CreateApp(root)
	app.Start()
}
