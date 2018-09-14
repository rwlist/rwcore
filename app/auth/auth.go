package auth

import (
	"net/http"
	"strings"
	"time"

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
	users       *Users
	userKey     string
	jwtSecret   []byte
	jwtDuration time.Duration
	jwtSigning  jwt.SigningMethod
}

func New(config Config) (*Auth, error) {
	duration, err := time.ParseDuration(config.JWT.Duration)
	if err != nil {
		return nil, err
	}
	return &Auth{
		users:       &Users{},
		userKey:     config.UserKey,
		jwtSecret:   []byte(config.JWT.Secret),
		jwtDuration: duration,
		jwtSigning:  jwt.GetSigningMethod(config.JWT.SigningMethod),
	}, nil
}

func (a *Auth) GetClaims(r *http.Request) (*Claims, error) {
	token := getToken(r)
	return a.readToken(token)
}

func getToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	token := strings.TrimPrefix(auth, BearerPrefix)
	return token
}
