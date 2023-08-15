package main

import (
	"Benim/api/route"
	"Benim/bootstrap"
	"net/http"
	"time"

	_ "Benim/assets/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Benim-Go
// @version 1.0
// @description This is a simple application to manage personal interests
// @termsOfService http://swagger.io/terms/

// @contact.name Saeed Sharghi Ghazani
// @contact.url https://github.com/saeedshargi
// @contact.email s.shargi.gh@gmail.com

// @license.name MIT
// @license.url https://opensource.org/license/mit/

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDB()

	cache := app.Redis
	defer app.CloseRedis()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	echo := echo.New()

	route.Setup(env, timeout, db, echo, cache)

	echo.GET("/", HealthCheck)
	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	echo.Logger.Fatal(echo.Start(env.ServerAddress))
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
