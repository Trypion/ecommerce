package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        string         `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    string         `gorm:"column:user_id;not null" json:"user_id"`
	Status    string         `gorm:"not null;default:'pending'" json:"status"`
	Total     float64        `gorm:"type:decimal(10,2);not null" json:"total"`
	Items     []OrderItem    `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"items"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   string    `gorm:"column:order_id;type:uuid;not null" json:"order_id"`
	ProductID string    `gorm:"column:product_id;not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
