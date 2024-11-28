package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"password,omitempty" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
