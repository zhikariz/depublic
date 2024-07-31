package common

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
