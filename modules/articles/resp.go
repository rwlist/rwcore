package articles

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/app/model"
	"github.com/rwlist/rwcore/app/utils"
)

type resp struct {
	impl impl
}

func (z resp) article(r *http.Request) model.Article {
	return r.Context().Value(articleKey).(model.Article)
}

func (z resp) get(w http.ResponseWriter, r *http.Request) {
	article := z.article(r)
	render.Respond(w, r, article)
}

func (z resp) getAll(w http.ResponseWriter, r *http.Request) {
	all, err := z.impl.getAll(r)
	utils.QuickRespond(w, r, all, err)
}

func (z resp) onClick(w http.ResponseWriter, r *http.Request) {
	res, err := z.impl.onClick(r, z.article(r))
	utils.QuickRespond(w, r, res, err)
}

func (z resp) setReadStatus(w http.ResponseWriter, r *http.Request) {
	res, err := z.impl.setReadStatus(r, z.article(r))
	utils.QuickRespond(w, r, res, err)
}

// func (m *Module) addOne(w http.ResponseWriter, r *http.Request) {
// 	var article model.Article
// 	err := render.Decode(r, &article)
// 	if err != nil {
// 		render.Render(w, r, utils.ErrBadRequest.With(err))
// 		return
// 	}
// 	article, err = addOneArticle(r, article)
// 	if err != nil {
// 		render.Render(w, r, utils.ErrInternal.With(err))
// 		return
// 	}
// 	render.Respond(w, r, article)
// }

// func (m *Module) addMany(w http.ResponseWriter, r *http.Request) {
// 	var articles []model.Article
// 	err := render.Decode(r, &articles)
// 	if err != nil {
// 		render.Render(w, r, utils.ErrBadRequest.With(err))
// 		return
// 	}
// 	articles, err = addManyArticles(r, articles)
// 	if err != nil {
// 		render.Render(w, r, utils.ErrInternal.With(err))
// 		return
// 	}
// 	render.Respond(w, r, articles)
// }

// func (m *Module) getAll(w http.ResponseWriter, r *http.Request) {
// 	articles, err := getAllArticles(r)
// 	if err != nil {
// 		render.Render(w, r, utils.ErrInternal.With(err))
// 	}
// 	render.Respond(w, r, articles)
// }

// func (m *Module) patch(w http.ResponseWriter, r *http.Request) {
// 	// TODO: parse {id} from link
// 	var article model.Article
// 	err := render.Decode(r, &article)
// 	if err != nil {
// 		render.Render(w, r, utils.ErrBadRequest.With(err))
// 		return
// 	}
// 	spew.Dump("Patch", article)
// 	article, err = patchArticle(r, article)
// 	if err != nil {
// 		render.Render(w, r, utils.ErrInternal.With(err))
// 		return
// 	}
// 	spew.Dump("Patch result", article)
// 	render.Respond(w, r, article)
// }