package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/wire"
	"github.com/rwlist/rwcore/users"
)

type JWTConfig struct {
	Secret        string
	Duration      string
	SigningMethod string
}

type Claims struct {
	User users.User
	jwt.StandardClaims
}

type LoginForm struct {
	Username string
	Password string
}

type SignUpForm struct {
	Username   string
	Email      string
	Password   string
	FirstName  string
	SecondName string
}

type TokenResponse struct {
	User  users.User
	Token string
}

var All = wire.NewSet(
	NewAuth,
	NewStdClaims,
	NewController,
	NewJWT,
	NewMiddleware,
	NewRouter,
)