package models

import "time"

type file struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	OwnerID    uint      `json:"owner_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SharedWith []uint    `json:"shared_with" gorm:"-"`
	FileType   string    `json:"file_type"`
	GitHistory []string  `json:"git_history" gorm:"-"`
}

type FileShare struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	FileID    uint      `json:"file_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GitCommit struct {
	ID uint `json:"id" gorm:"primarykey"`
	FileID uint `json:"file_id"`
	Hash string `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
}