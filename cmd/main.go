package main

import (
	"Benim/api/route"
	"Benim/bootstrap"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDB()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	echo := echo.New()

	route.Setup(env, timeout, db, echo)

	echo.Logger.Fatal(echo.Start(env.ServerAddress))
}
