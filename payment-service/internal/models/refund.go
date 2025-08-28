package models

import (
	"time"

	"gorm.io/gorm"
)

type Refund struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	PaymentID string `gorm:"not null;index" json:"payment_id"`

	Amount   float64      `gorm:"not null;type:decimal(10,2)" json:"amount"`
	Currency string       `gorm:"not null;size:3" json:"currency"`
	Reason   string       `gorm:"size:500" json:"reason"`
	Status   RefundStatus `gorm:"not null;default:'pending'" json:"status"`

	// Provider details
	ProviderID       string `gorm:"size:255" json:"provider_id"`
	ProviderResponse string `gorm:"type:text" json:"-"`

	// Admin info
	ProcessedBy string `gorm:"size:255" json:"processed_by"` // Admin user ID

	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ProcessedAt *time.Time     `json:"processed_at,omitempty"`

	// Relationship
	Payment Payment `gorm:"foreignKey:PaymentID" json:"payment,omitempty"`
}

type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusCompleted RefundStatus = "completed"
	RefundStatusFailed    RefundStatus = "failed"
	RefundStatusCancelled RefundStatus = "cancelled"
)
