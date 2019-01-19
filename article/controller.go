package article

import (
	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/resp"
	"net/http"
	"strconv"
)

type Controller struct {
	Service Service
}

func NewController(s Service) *Controller {
	return &Controller{
		Service: s,
	}
}

func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	article := R(r)
	render.Respond(w, r, article)
}

func (c *Controller) addURL(w http.ResponseWriter, r *http.Request) {
	db := DB(r)

	var url string
	err := render.Decode(r, &url)
	if err != nil {
		render.Render(w, r, resp.ErrBadRequest.With(err))
		return
	}
	res, err := c.Service.AddURL(db, url)
	resp.QuickRespond(w, r, res, err)
}

func (c *Controller) getAll(w http.ResponseWriter, r *http.Request) {
	db := DB(r)

	all, err := db.GetAll()
	for i := len(all)/2 - 1; i >= 0; i-- {
		opp := len(all) - 1 - i
		all[i], all[opp] = all[opp], all[i]
	}
	resp.QuickRespond(w, r, all, err)
}

func (c *Controller) onClick(w http.ResponseWriter, r *http.Request) {
	db := DB(r)
	article := R(r)

	res, err := c.Service.OnClick(db, article)
	resp.QuickRespond(w, r, res, err)
}

func (c *Controller) setReadStatus(w http.ResponseWriter, r *http.Request) {
	db := DB(r)
	article := R(r)

	newStatus := r.URL.Query().Get("newStatus")
	res, err := c.Service.SetReadStatus(db, article, newStatus)
	resp.QuickRespond(w, r, res, err)
}

func (c *Controller) changeRating(w http.ResponseWriter, r *http.Request) {
	db := DB(r)
	article := R(r)

	delta, err := strconv.Atoi(r.URL.Query().Get("delta"))
	if err != nil {
		render.Render(w, r, resp.ErrBadRequest.With(err))
		return
	}

	res, err := c.Service.ChangeRating(db, article, delta)
	resp.QuickRespond(w, r, res, err)
}

func (c *Controller) removeTag(w http.ResponseWriter, r *http.Request) {
	db := DB(r)
	article := R(r)
	tag := r.URL.Query().Get("tag")

	res, err := c.Service.RemoveTag(db, article, tag)
	resp.QuickRespond(w, r, res, err)
}

func (c *Controller) addTag(w http.ResponseWriter, r *http.Request) {
	db := DB(r)
	article := R(r)
	tag := r.URL.Query().Get("tag")

	res, err := c.Service.AddTag(db, article, tag)
	resp.QuickRespond(w, r, res, err)
}
