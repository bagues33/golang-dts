package models

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"primary_key"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	Message   string    `json:"message" gorm:"not null" validate:"required,min:1,max:255`
	CreatedAt time.Time `json:"created_at"` // created_at (Date)
	UpdatedAt time.Time `json:"updated_at"` // updated_at (Date)
	User      User      `json:"user" gorm:"foreignkey:UserID"`
	Photo     Photo     `json:"photo" gorm:"foreignkey:PhotoID"`
}
