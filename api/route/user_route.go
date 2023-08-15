package route

import (
	"Benim/api/controller"
	"Benim/bootstrap"
	"Benim/domain"
	"Benim/mongo"
	"Benim/redis"
	"Benim/repository"
	"Benim/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

func NewUserRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *echo.Group, cache redis.Client) {
	ur := repository.NewUserRepository(db, domain.CollectionUsers)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		Secret:      env.JwtSecretkey,
	}

	group.POST("", uc.Signup)
}
