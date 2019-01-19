package auth

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrTokenInvalid = errors.New("JWT token invalid")
	ErrBadSigningMethod = errors.New("unexpected signing method")
)

type JWT struct{
	secret []byte
	signingMethod jwt.SigningMethod
}

func NewJWT(c JWTConfig) *JWT {
	return &JWT{
		secret: []byte(c.Secret),
		signingMethod: jwt.GetSigningMethod(c.SigningMethod),
	}
}

func (j *JWT) CreateAndSign(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(j.signingMethod, claims)
	tokenString, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// TODO: verify expiration?
func (j *JWT) ParseAndVerify(value string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(value, claims, j.keyFunc)
	if err != nil {
		return err
	}
	if !token.Valid {
		return ErrTokenInvalid
	}
	return nil
}


func (j *JWT) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, ErrBadSigningMethod
	}
	return j.secret, nil
}