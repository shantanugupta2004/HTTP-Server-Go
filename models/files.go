package models

import(
	"time"
)

type File struct{
	ID uint `gorm:"primaryKey"`
	FileName string
	FilePath string
	FileSize int64
	UploadedAt time.Time
	UserID uint
}