package models

import(
	"time"
)

type File struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	FileName   string    `json:"fileName"`
	FilePath   string    `json:"filePath"`
	FileSize   int64     `json:"fileSize"`
	UploadedAt time.Time `json:"uploadedAt"`
	IsShared   bool      `json:"isShared"`
	ShareToken string    `json:"shareToken"`
	UserID     uint      `json:"userID"`
}
