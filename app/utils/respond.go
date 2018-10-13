package utils

import "net/http"

func customRespond(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.DefaultResponser(w, r, v)
}

func init() {
	render.Respond = customRespond
}
