package auth

import (
	"net/http"
)

func (a *Auth) handleLogin(r *http.Request) (*Claims, error) {
	user, err := a.users.FindUser(r)
	if err != nil {
		return nil, err
	}
	return a.newClaims(user)
}

func (a *Auth) handleSignUp(r *http.Request) (*Claims, error) {
	user, err := a.users.CreateUser(r)
	if err != nil {
		return nil, err
	}
	return a.newClaims(user)
}
