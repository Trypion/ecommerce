package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey;type:uuid;default:uuidv7()" json:"id"`
	Email    string `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password string `gorm:"not null;size:255" json:"-"`
	Name     string `gorm:"not null;size:100" json:"name"`
	Role     string `gorm:"not null;size:20" json:"role"` // customer, admin, etc.

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
