package db

import (
	"context"
	"github.com/kletskovg/typecode/server/db_service/src/internal/utils"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type File struct {
	Name string
	Code [][]string
	Owner string
	Language string
}

func Connect() *mongo.Client {
	
	DB_CONNECTION := utils.GetEnvVar("DB_CONNECTION")

	log.Info("Starting DB connection")
	log.Info(DB_CONNECTION)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DB_CONNECTION))
	if err != nil { log.Error(err) }

	log.Info("Successfully connected")
	log.Info(client)

	go clearCache(client)
	return client
}
