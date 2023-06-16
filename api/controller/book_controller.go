package controller

import (
	"Benim/domain"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	BookUsecase domain.BookUsecase
}

var validate = validator.New()

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the input paylod
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body domain.Book true "Create book"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /books [post]
func (bc *BookController) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book domain.Book
	defer cancel()

	//validate the request body
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&book); validationErr != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: validationErr.Error()})
	}

	err := bc.BookUsecase.Create(ctx, &book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Book created successfully."})
}
