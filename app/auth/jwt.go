package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func (a *Auth) createToken(claims jwt.MapClaims) (string, error) {
	cur := time.Now().UTC()
	exp := cur.Add(a.jwtDuration)
	claims["nbf"] = cur.Unix()
	claims["exp"] = exp.Unix()
	token := jwt.NewWithClaims(a.jwtSigning, claims)
	tokenString, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) readToken(value string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return a.jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("JWT token invalid")
}
