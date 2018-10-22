package articles

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/app/auth"
	"github.com/rwlist/rwcore/app/model"
	"github.com/rwlist/rwcore/app/utils"
)

func (m *Module) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(auth.HasRole("articles"))

	r.Post("/add", m.addOne)
	r.Post("/addMany", m.addMany)
	r.Get("/all", m.getAll)
	r.Route("/{id}", func(r chi.Router) {
		r.Patch("/patch", m.patch)
	})

	return r
}

func (m *Module) addOne(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	err := render.Decode(r, &article)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest.With(err))
		return
	}
	article, err = addOneArticle(r, article)
	if err != nil {
		render.Render(w, r, utils.ErrInternal.With(err))
		return
	}
	render.Respond(w, r, article)
}

func (m *Module) addMany(w http.ResponseWriter, r *http.Request) {
	var articles []model.Article
	err := render.Decode(r, &articles)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest.With(err))
		return
	}
	articles, err = addManyArticles(r, articles)
	if err != nil {
		render.Render(w, r, utils.ErrInternal.With(err))
		return
	}
	render.Respond(w, r, articles)
}

func (m *Module) getAll(w http.ResponseWriter, r *http.Request) {
	articles, err := getAllArticles(r)
	if err != nil {
		render.Render(w, r, utils.ErrInternal.With(err))
	}
	render.Respond(w, r, articles)
}

func (m *Module) patch(w http.ResponseWriter, r *http.Request) {
	// TODO: parse {id} from link
	var article model.Article
	err := render.Decode(r, &article)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest.With(err))
		return
	}
	article, err = patchArticle(r, article)
	if err != nil {
		render.Render(w, r, utils.ErrInternal.With(err))
		return
	}
	render.Respond(w, r, article)
}
