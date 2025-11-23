package database

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Database, error) {
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return nil, err
    }

    return client.Database("uas_be_mongo"), nil
}
