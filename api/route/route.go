package route

import (
	"Benim/bootstrap"
	"Benim/mongo"
	"Benim/redis"
	"time"

	"github.com/labstack/echo/v4"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, echo *echo.Echo, cache redis.Client) {
	publicRoute := echo.Group("/users")
	NewUserRoute(env, timeout, db, publicRoute, cache)

	protectedRoute := echo.Group("/books")
	// protectedRoute.Use(middleware.JwtAuthenticationMiddleware(env.JwtSecretkey))
	NewBookRoute(env, timeout, db, protectedRoute, cache)

	protectedLanguagesRoute := echo.Group("/languages")
	// protectedRoute.Use(middleware.JwtAuthenticationMiddleware(env.JwtSecretkey))
	NewLanguageRoute(env, timeout, db, protectedLanguagesRoute, cache)
}
