package auth

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

func (a *Auth) createToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(a.jwtSigning, claims)
	tokenString, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("Unexpected signing method")
	}
	return a.jwtSecret, nil
}

func (a *Auth) readToken(value string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(value, claims, a.keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("JWT token invalid")
	}
	return claims, nil
}
