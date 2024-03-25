package models

import "time"

type SocialMedia struct {
	ID             int       `json:"id"`
	Name           string    `json:"name" gorm:"not null" validate:"required,min=3"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" validate:"required,url,min=3"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `json:"user" gorm:"foreignkey:UserID"`
}
