package auth

import "net/http"

type Provider interface {
	FindUser(r *http.Request) (interface{}, error)
	CreateUser(r *http.Request) (interface{}, error)
}
