package database

import (
	"context"
	"log"
	"time"

	"projectuasbe/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func InitMongo() {
	cfg := config.AppConfig

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("❌ Failed to create Mongo client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to connect MongoDB: %v", err)
	}

	MongoDB = client.Database(cfg.MongoDB)
	log.Println("✅ MongoDB connected successfully")
}
