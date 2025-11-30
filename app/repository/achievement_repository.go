package repository

import (
	"context"
	"errors"
	"projectuasbe/app/model"
	"projectuasbe/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AchievementRepository interface {
	Create(ctx context.Context, a *model.Achievement) (*mongo.InsertOneResult, error)
	FindByID(ctx context.Context, id string) (*model.Achievement, error)
	FindByStudent(ctx context.Context, studentID string) ([]model.Achievement, error)
	Update(ctx context.Context, id string, update bson.M) error
	SoftDelete(ctx context.Context, id string) error
}

type achievementRepository struct {
	col *mongo.Collection
}

func NewAchievementRepository() AchievementRepository {
	return &achievementRepository{
		col: database.MongoDB.Collection("achievements"),
	}
}

//
// ==============================
//            CREATE
// ==============================
//

func (r *achievementRepository) Create(ctx context.Context, a *model.Achievement) (*mongo.InsertOneResult, error) {
	a.ID = primitive.NewObjectID().Hex()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	a.Status = "draft"
	a.DeletedAt = nil

	return r.col.InsertOne(ctx, a)
}

//
// ==============================
//            FIND BY ID
// ==============================
//

func (r *achievementRepository) FindByID(ctx context.Context, id string) (*model.Achievement, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var result model.Achievement
	err = r.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

//
// ==============================
//       FIND BY STUDENT ID
// ==============================
//

func (r *achievementRepository) FindByStudent(ctx context.Context, studentID string) ([]model.Achievement, error) {
	cursor, err := r.col.Find(ctx, bson.M{
		"studentId": studentID,
		"status": bson.M{"$ne": "deleted"},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []model.Achievement
	err = cursor.All(ctx, &results)
	return results, err
}

//
// ==============================
//             UPDATE
// ==============================
//

func (r *achievementRepository) Update(ctx context.Context, id string, update bson.M) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	update["updatedAt"] = time.Now()

	_, err = r.col.UpdateByID(ctx, objID, bson.M{
		"$set": update,
	})
	return err
}

//
// ==============================
//          SOFT DELETE
// ==============================
//

func (r *achievementRepository) SoftDelete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	now := time.Now()

	_, err = r.col.UpdateByID(ctx, objID, bson.M{
		"$set": bson.M{
			"status":    "deleted",
			"deletedAt": now,
			"updatedAt": now,
		},
	})

	return err
}
