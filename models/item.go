package models

type Item struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	ItemCode    string `json:"item_code" gorm:"unique;not null"`
	Description string `json:"description" gorm:"not null"`
	Quantity    uint   `json:"quantity" gorm:"not null"`
	OrderId     uint   `json:"order_id"`
}
