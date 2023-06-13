package bootstrap

import (
	"Benim/mongo"

	"context"
	"fmt"
	"log"
	"time"
)

func ConnectDB(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin&readPreference=primary&ssl=false&directConnection=true", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseDB(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

func InitDatabase(client mongo.Client, databaseName string) {
	if client == nil {
		log.Fatal("Error in connection to database.")
	}

	err := client.Database(databaseName)
	if err != nil {
		log.Fatal("Error on creating document")
	}

	log.Panicf("Document %s created successfully.", databaseName)
}
