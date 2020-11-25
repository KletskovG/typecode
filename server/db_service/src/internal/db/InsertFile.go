package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	log "github.com/sirupsen/logrus"
	"context"
)


func InsertFile(client *mongo.Client, file File) *mongo.InsertOneResult {
	db := client.Database("Gym")
	cache := db.Collection("cache", nil)

	res, err := cache.InsertOne(context.TODO(), bson.M{
		"name": file.Name,
		"code": file.Code,
		"owner": file.Owner,
		"language": file.Language,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"file": file,
		}).Error("During inserting file to cache")
	}

	log.WithFields(log.Fields{
		"file": file,
		"res": res,
	}).Info("Succefully save file to cache")

	return res
}