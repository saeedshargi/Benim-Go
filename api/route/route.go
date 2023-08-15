package route

import (
	"Benim/api/middleware"
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
	protectedRoute.Use(middleware.JwtAuthenticationMiddleware(env.JwtSecretkey))
	NewBookRoute(env, timeout, db, protectedRoute, cache)
}
