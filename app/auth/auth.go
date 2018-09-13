package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/rwlist/rwcore/app/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	BearerPrefix = "Bearer "
)

type Config struct {
	JWT     JWTConfig
	UserKey string
}

type JWTConfig struct {
	Secret        string
	Duration      string
	SigningMethod string
}

type Auth struct {
	provider    Provider
	userKey     string
	jwtSecret   []byte
	jwtDuration time.Duration
	jwtSigning  jwt.SigningMethod
}

func New(provider Provider, config Config) (*Auth, error) {
	duration, err := time.ParseDuration(config.JWT.Duration)
	if err != nil {
		return nil, err
	}
	return &Auth{
		provider:    provider,
		userKey:     config.UserKey,
		jwtSecret:   []byte(config.JWT.Secret),
		jwtDuration: duration,
		jwtSigning:  jwt.GetSigningMethod(config.JWT.SigningMethod),
	}, nil
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
	user := claims[a.userKey]
	if user == nil {
		return nil, errors.New("No user in claims")
	}
	return user, nil
}

func (a *Auth) processUser(w http.ResponseWriter, r *http.Request, user interface{}, err error) {
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest.With(err))
		return
	}
	token, err := a.createToken(jwt.MapClaims{
		a.userKey: user,
	})
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest.With(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, TokenResponse{user, token})
}

func getToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	token := strings.TrimPrefix(auth, BearerPrefix)
	return token
}
