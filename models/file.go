package models

import (
	"time"

	"gorm.io/gorm"
)

type IntArray []int

type file struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	OwnerID    uint      `json:"owner_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SharedWith []uint    `json:"shared_with" gorm:"type:json"`
	FileType   string    `json:"file_type"`
	GitHistory []string  `json:"git_history" gorm:"type-json"`
}

func (f *File) BeforeCreate(tx *gorm.DB) error {
	f.FileType = "owned"
	return nil
}
