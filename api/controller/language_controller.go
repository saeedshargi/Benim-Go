package controller

import (
	"Benim/domain"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LanguageController struct {
	LanguageUsecase domain.LanguageUsecase
}

// GetAllLanguage godoc
// @Summary Get all languages
// @Description Get all language with the input paylod
// @Tags languages
// @Accept  json
// @Produce  json
// @Param language body domain.Language true "Get All language"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /languages [post]
func (bc *LanguageController) GetAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	languages, err := bc.LanguageUsecase.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Language fetched successfully.", Data: languages})
}

// CreateLanguage godoc
// @Summary Create a new language
// @Description Create a new language with the input paylod
// @Tags languages
// @Accept  json
// @Produce  json
// @Param language body domain.Language true "Create language"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /languages [post]
func (bc *LanguageController) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var language domain.Language
	defer cancel()

	//validate the request body
	if err := c.Bind(&language); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&language); validationErr != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: validationErr.Error()})
	}

	err := bc.LanguageUsecase.Create(ctx, &language)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Language created successfully."})
}
