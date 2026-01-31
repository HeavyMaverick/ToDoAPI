package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          int            `json:"id" gorm:"primary_key;autoIncrement"`
	Title       string         `json:"title" gorm:"not null,size:255"`
	Description string         `json:"description" gorm:"type:text"`
	Completed   bool           `json:"completed" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	UserID      int            `json:"user_id" gorm:"not null;default:1"`
}
