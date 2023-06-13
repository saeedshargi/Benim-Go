package bootstrap

import "Benim/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = LoadEnv()
	app.Mongo = ConnectDB(app.Env)
	InitDatabase(app.Mongo, app.Env.DBName)
	return *app
}

func (app *Application) CloseDB() {
	CloseDB(app.Mongo)
}
