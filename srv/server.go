package srv

import (
	"github.com/google/wire"
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

func New(c Config, r router.Router) *Server {
	srv := &Server{
		Router:   r,
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

var All = wire.NewSet(
	New,
)