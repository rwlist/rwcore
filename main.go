package main

import (
	"flag"
	"log"

	_ "github.com/rwlist/rwcore/app/resp"
	"github.com/rwlist/rwcore/srv"
)

type App struct {
	Server *srv.Server
}

func main() {
	log.SetFlags(log.Lshortfile)

	filepath := *flag.String("config", "conf/config.toml", "pass path to config file")
	flag.Parse()

	b, err := Initialize(filepath)
	if err != nil {
		log.Println("Initialization failed.", err)
		return
	}

	b.Server.Start()
}
