package repository

import (
    "context"
    "projectuasbe/app/model"
    "projectuasbe/database"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type AchievementRepository interface {
    GetByID(ctx context.Context, id string) (*model.Achievement, error)
    Insert(ctx context.Context, a *model.Achievement) error
}

type achievementRepository struct {
    collection *mongo.Collection
}

func NewAchievementRepository() AchievementRepository {
    return &achievementRepository{
        collection: database.MongoDB.Collection("achievements"),
    }
}

func (r *achievementRepository) GetByID(ctx context.Context, id string) (*model.Achievement, error) {
    var a model.Achievement

    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&a)
    if err != nil {
        return nil, err
    }

    return &a, nil
}

func (r *achievementRepository) Insert(ctx context.Context, a *model.Achievement) error {
    _, err := r.collection.InsertOne(ctx, a)
    return err
}
