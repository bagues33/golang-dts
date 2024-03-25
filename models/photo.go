package models

import "time"

type Photo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"not null" validate:"required,min:1,max:255`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" gorm:"not null" validate:"required,url`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignkey:UserID"`
	Comments  []Comment
}
