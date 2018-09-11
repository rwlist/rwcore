package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/rwlist/rwcore/app/utils"
)

const (
	BearerPrefix = "Bearer "
	UserKey      = "User"
)

type Auth struct {
	provider    Provider
	jwtSecret   []byte
	jwtDuration time.Duration
	jwtSigning  jwt.SigningMethod
}

func New(provider Provider) *Auth {
	return &Auth{
		provider:  provider,
		jwtSecret: []byte("CHANGE_THIS"), // TODO
	}
}

func (a *Auth) GetClaims(r *http.Request) (jwt.MapClaims, error) {
	token := getToken(r)
	return a.readToken(token)
}

func (a *Auth) GetUser(r *http.Request) (interface{}, error) {
	claims, err := a.GetClaims(r)
	if err != nil {
		return nil, err
	}
	user := claims[UserKey]
	if user == nil {
		return nil, errors.New("No user in claims")
	}
	return user, nil
}

func (a *Auth) processUser(w http.ResponseWriter, user interface{}, err error) {
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	token, err := a.createToken(jwt.MapClaims{
		UserKey: user,
	})
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteJSON(w, http.StatusOK, TokenResponse{user, token})
}

func getToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	token := strings.TrimPrefix(auth, BearerPrefix)
	return token
}
