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

func NewLanguageRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *echo.Group, cache redis.Client) {
	lr := repository.NewLanguageRepository(db, domain.CollectionLanguage)
	lc := &controller.LanguageController{
		LanguageUsecase: usecase.NewLanguageUsecase(lr, timeout),
	}

	group.POST("", lc.Create)

}
