package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type StdClaims struct {
	duration time.Duration
}

func NewStdClaims(c *JWTConfig) (*StdClaims, error) {
	duration, err := time.ParseDuration(c.Duration)
	if err != nil {
		return nil, err
	}
	return &StdClaims{
		duration: duration,
	}, nil
}

func (c StdClaims) Create() jwt.StandardClaims {
	cur := time.Now().UTC()
	exp := cur.Add(c.duration)
	return jwt.StandardClaims{
		NotBefore: cur.Unix(),
		IssuedAt:  cur.Unix(),
		ExpiresAt: exp.Unix(),
	}
}
