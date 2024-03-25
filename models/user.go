package models

import (
	"time"
	// "github.com/go-playground/validator/v10"
)

type User struct {
	ID           int    // Primary key
	Username     string `gorm:"unique;not null" validate:"required,gte=3,lte=20"`
	Email        string `gorm:"unique;not null email" validate:"required,email"`
	Password     string `gorm:"not null;min:6" validate:"required,min=6"`
	Age          int    `gorm:"not null;min:8" validate:"required,min=8"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Photos       []Photo
	SocialMedias []SocialMedia
	Comments     []Comment
}
