package controller

import (
	"Benim/domain"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Secret      string
}

// Signup godoc
// @Summary Signup a new user
// @Description Signup a new user with the input paylod
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body domain.User true "Signup"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /users/signup [post]
func (uc *UserController) Signup(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user domain.SignupUser
	defer cancel()

	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: validationErr.Error()})
	}

	userResponse, err := uc.UserUsecase.Register(ctx, &user, uc.Secret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{Message: "User created successfully.", Data: userResponse})
}
