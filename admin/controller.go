package admin

import (
	"net/http"

	"github.com/go-chi/render"
)

type Controller struct{}

func (s Controller) test(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Message{"Congrats, you are admin!"})
}

func NewController() Controller {
	return Controller{}
}
