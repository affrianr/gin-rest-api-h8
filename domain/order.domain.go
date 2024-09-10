package domain

import "time"

type Order struct {
	ID           uint      `json:"order_id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}
