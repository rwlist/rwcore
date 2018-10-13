package lists

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (l *Lists) Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", l.index)
	r.Route("/{name}", func(r chi.Router) {
		r.Get("/data", l.data)
		r.Post("/insertOne", l.insertOne)
		r.Post("/insertMany", l.insertMany)
		r.Post("/clear", l.clear)
		r.Get("/backup", l.backup)
		r.Post("/copyToDir", l.copyToDir)
	})

	return r
}

func (l *Lists) index(w http.ResponseWriter, r *http.Request) {

}

func (l *Lists) data(w http.ResponseWriter, r *http.Request) {

}

func (l *Lists) insertOne(w http.ResponseWriter, r *http.Request) {

}

func (l *Lists) insertMany(w http.ResponseWriter, r *http.Request) {

}

func (l *Lists) clear(w http.ResponseWriter, r *http.Request) {

}

func (l *Lists) backup(w http.ResponseWriter, r *http.Request) {

}

func (l *Lists) copyToDir(w http.ResponseWriter, r *http.Request) {

}
