package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

func customRespond(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.DefaultResponder(w, r, v)
}

func init() {
	render.Respond = customRespond
}
