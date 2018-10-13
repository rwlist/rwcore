package articles

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/app/utils"
)

func (a *Articles) Router() chi.Router {
	r := chi.NewRouter()

	r.Post("/add", a.addOne)
	r.Post("/addMany", a.addMany)
	r.Get("/all", a.getAll)
	r.Route("/one/{id}", func(r chi.Router) {
		r.Patch("/status", l.patchStatus)
		r.Patch("/tags", l.patchTags)
		r.Patch("/name", l.patchName)
	})

	return r
}

func (a *Articles) addOne(w http.ResponseWriter, r *http.Request) {

}

func (a *Articles) addMany(w http.ResponseWriter, r *http.Request) {

}

func (a *Articles) getAll(w http.ResponseWriter, r *http.Request) {
	articles, err := getAllArticles(r)
	if err != nil {
		render.Render(w, r, utils.ErrInternal.With(err))
	}
	render.JSON(w, r, articles)
}

func (a *Articles) patchStatus(w http.ResponseWriter, r *http.Request) {
	render.Bind()
	article, err := patchStatus()
}

func (a *Articles) patchTags(w http.ResponseWriter, r *http.Request) {

}

func (a *Articles) patchName(w http.ResponseWriter, r *http.Request) {

}
