package resp

import (
	"net/http"

	"github.com/go-chi/render"
)

func customRespond(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.DefaultResponder(w, r, v)
}

func QuickRespond(w http.ResponseWriter, r *http.Request, v interface{}, err error) {
	if err != nil {
		render.Render(w, r, ErrInternal.With(err))
		return
	}
	render.Respond(w, r, v)
}

func init() {
	render.Respond = customRespond
}
