package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"context"
	"time"
	log "github.com/sirupsen/logrus"
)

var instance *mongo.Client

var DB_CONNECTION = os.Getenv("DB_CONNECTION")

func Connect() *mongo.Client {
	if instance == nil {
		log.Info("Starting DB connection")
		ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(DB_CONNECTION))
		if err != nil { log.Fatal(err) }

		instance = client
		log.Info("Successfully connected")
		log.Info(instance)
	}

	return instance
}