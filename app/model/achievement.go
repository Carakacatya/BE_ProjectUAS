package model

import "time"

type Achievement struct {
    ID            string                 `json:"id"`
    StudentID     string                 `json:"student_id"`
    AchievementType string               `json:"achievement_type"`
    Title         string                 `json:"title"`
    Description   string                 `json:"description"`
    Details       map[string]interface{} `json:"details"`
    Attachments   []Attachment           `json:"attachments"`
    Tags          []string               `json:"tags"`
    Points        float64                `json:"points"`

    Status        string                 `json:"status"`     // added
    DeletedAt     *time.Time             `json:"deleted_at"` // added for soft delete

    CreatedAt     time.Time              `json:"created_at"`
    UpdatedAt     time.Time              `json:"updated_at"`
}

type Attachment struct {
    FileName   string    `json:"file_name"`
    FileURL    string    `json:"file_url"`
    FileType   string    `json:"file_type"`
    UploadedAt time.Time `json:"uploaded_at"`
}
