package domain

import "time"

type Item struct {
	ID          uint      `json:"item_id" gorm:"primaryKey"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     uint      `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

