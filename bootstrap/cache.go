package bootstrap

import (
	"Benim/redis"
	"context"
	"fmt"
	"log"
	"time"
)

func ConnectRedis(env *Env) redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	redisAddress := env.RedisAddress
	redisPort := env.RedisPort

	redisUrl := fmt.Sprintf("%s:%s", redisAddress, redisPort)

	client := redis.NewClient(redisUrl)

	err := client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseRedis(client redis.Client) {
	if client == nil {
		return
	}

	err := client.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to Redis closed.")
}
