package boot

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Url      string
	Database string
}

type MongoApp struct {
	config *MongoConfig
	db     *mongo.Database
}

func NewMongoApp(config *MongoConfig) *MongoApp {
	return &MongoApp{
		config: config,
	}
}

func (s *MongoApp) Run(done chan error) {
	clientOptions := options.Client().ApplyURI(s.config.Url)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		done <- err
		log.Panicf("Error on connect mongo: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		done <- err
		log.Panicf("Error on ping mongo: %v", err)
	}

	db := client.Database(s.config.Database)
	_ = db

	s.db = db
	log.Println("Mongo Connected")
	done <- nil
}

func (s MongoApp) GetDb() *mongo.Database {
	return s.db
}
