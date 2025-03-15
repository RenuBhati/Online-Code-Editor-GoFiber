package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (ia IntArray) Value() (driver.Value, error) {
	return json.Marshal(ia)
}

func (ia *IntArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &ia)
}

type StringArray []StringArray

func (sa StringArray) Value() (driver.Value, error) {
	return json.Marshal(sa)
}

func (sa *StringArray) Scan(value interface{}) error {
	bytes, ok := Value([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &sa)
}

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
