package route

import (
	"Benim/api/controller"
	"Benim/bootstrap"
	"Benim/domain"
	"Benim/mongo"
	"Benim/repository"
	"Benim/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

func NewBookRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *echo.Group) {
	br := repository.NewBookRepository(db, domain.CollectionBook)
	bc := &controller.BookController{
		BookUsecase: usecase.NewBookUsecase(br, timeout),
	}

	group.POST("", bc.Create)
}
