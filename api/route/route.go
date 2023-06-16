package route

import (
	"Benim/bootstrap"
	"Benim/mongo"
	"time"

	"github.com/labstack/echo/v4"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, echo *echo.Echo) {
	protectedRoute := echo.Group("/books")
	NewBookRoute(env, timeout, db, protectedRoute)
}
