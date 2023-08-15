package domain

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Email    string `json:"email"`
	FullName string `json:"fulName"`
	Id       string `json:"id"`
	jwt.Claims
}

type JwtCustomRefreshClaims struct {
	Id string `json:"id"`
	jwt.Claims
}
