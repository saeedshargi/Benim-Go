package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	RedisAddress   string `mapstructure:"REDIS_ADDRESS"`
	RedisPort      string `mapstructure:"REDIS_PORT"`
	JwtSecretkey   string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadEnv() *Env {
	env := Env{}
	viper.AddConfigPath("../")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the .env file", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Can't load the .env file", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development enviroment.")
	}

	return &env
}
