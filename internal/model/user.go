package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"not null;unique;size:30"`
	Email     string    `json:"email" gorm:"not null;unique;size:60"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Task []Task `json:"tasks,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
