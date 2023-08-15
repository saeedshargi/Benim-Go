package tokenutil

import (
	"Benim/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *domain.User, secret string, expireTime int) (accessToken string, err error) {
	exp := time.Now().Add(time.Minute * time.Duration(expireTime)).UTC()
	claims := &domain.JwtCustomClaims{
		Email:    user.Email,
		FullName: user.FullName,
		Id:       user.ID.Hex(),
		Claims: jwt.MapClaims{
			"exp": exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	accToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return accToken, nil
}

func CreateRefreshToken(user *domain.User, secret string, expireTime int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expireTime)).UTC()
	claims := &domain.JwtCustomRefreshClaims{
		Id: user.ID.Hex(),
		Claims: jwt.MapClaims{
			"exp": exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	refToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return refToken, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIdFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims["id"].(string), nil
}
