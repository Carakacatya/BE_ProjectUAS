package model

import "time"

type Attachment struct {
	FileName   string    `bson:"fileName" json:"fileName"`
	FileUrl    string    `bson:"fileUrl" json:"fileUrl"`
	FileType   string    `bson:"fileType" json:"fileType"`
	UploadedAt time.Time `bson:"uploadedAt" json:"uploadedAt"`
}

type Achievement struct {
	ID              string      `bson:"_id,omitempty" json:"id"`
	StudentID       string      `bson:"studentId" json:"studentId"`
	AchievementType string      `bson:"achievementType" json:"achievementType"`
	Title           string      `bson:"title" json:"title"`
	Description     string      `bson:"description" json:"description"`
	Details         any         `bson:"details" json:"details"` // flexible field
	Attachments     []Attachment `bson:"attachments" json:"attachments"`
	Tags            []string    `bson:"tags" json:"tags"`
	Points          float64     `bson:"points" json:"points"`
	Status          string      `bson:"status" json:"status"` // draft, submitted, verified, rejected, deleted
	CreatedAt       time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time   `bson:"updatedAt" json:"updatedAt"`
	DeletedAt       *time.Time  `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}
