package bootstrap

import (
	"Benim/mongo"
	"Benim/redis"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Redis redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = LoadEnv()
	app.Mongo = ConnectDB(app.Env)
	app.Redis = ConnectRedis(app.Env)
	return *app
}

func (app *Application) CloseDB() {
	CloseDB(app.Mongo)
}

func (app *Application) CloseRedis() {
	CloseRedis(app.Redis)
}
