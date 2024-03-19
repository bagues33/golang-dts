package models

type Order struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	CustomerName string `json:"customer_name" gorm:"not null"`
	OrderedAt    string `json:"ordered_at" gorm:"not null"`
	Items        []Item
}
