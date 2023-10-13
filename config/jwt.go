package config

import (
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("efbefbbby74792b29br327f3b9834")

type JWTClaim struct {
	Username	string
	jwt.RegisteredClaims
}