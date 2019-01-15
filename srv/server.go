package srv

import (
	"log"
	"net/http"

	"github.com/rwlist/rwcore/router"
)

type Server struct {
	Router   router.Router
	BindAddr string
}

type Config struct {
	BindAddr string
}

type Deps struct {
	Router router.Router
}

func New(c Config, deps Deps) *Server {
	srv := &Server{
		Router:   deps.Router,
		BindAddr: c.BindAddr,
	}

	return srv
}

func (s *Server) Start() {
	log.Printf("Server started on %s\n", s.BindAddr)

	if err := http.ListenAndServe(s.BindAddr, s.Router); err != nil {
		log.Println("Server exited.", err)
	}
}
