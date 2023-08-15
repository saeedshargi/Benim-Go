package middleware

import (
	"Benim/domain"
	"Benim/internal/tokenutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func JwtAuthenticationMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			token := strings.Split(authHeader, " ")
			if len(token) == 2 {
				authToken := token[1]
				authorized, err := tokenutil.IsAuthorized(authToken, secret)
				if authorized {
					userId, err := tokenutil.ExtractIdFromToken(authToken, secret)
					if err != nil {
						c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
						c.Error(err)
						return echo.ErrUnauthorized
					}
					c.Set("x-user-id", userId)
					return next(c)
				}
				c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
				c.Error(err)
				return echo.ErrUnauthorized
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
			return echo.ErrUnauthorized
		}
	}
}
