// payment-service/internal/models/payment.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID      string `gorm:"primaryKey;type:uuid;default:uuidv7()" json:"id"`
	OrderID string `gorm:"not null;index" json:"order_id"`
	UserID  string `gorm:"not null;index" json:"user_id"`

	// Amount details
	Amount   float64 `gorm:"not null;type:decimal(10,2)" json:"amount"`
	Currency string  `gorm:"not null;default:'USD';size:3" json:"currency"`

	// Payment status and method
	Status PaymentStatus `gorm:"not null;default:'pending'" json:"status"`
	Method PaymentMethod `gorm:"not null" json:"method"`

	// Provider details
	Provider         string `gorm:"size:50" json:"provider"`           // stripe, paypal, square
	ProviderID       string `gorm:"size:255;index" json:"provider_id"` // External payment ID
	ProviderResponse string `gorm:"type:text" json:"-"`                // Raw provider response

	// Refund information
	RefundedAmount float64 `gorm:"type:decimal(10,2);default:0" json:"refunded_amount"`
	IsRefunded     bool    `gorm:"default:false" json:"is_refunded"`

	// Failure information
	FailureCode    string `gorm:"size:50" json:"failure_code,omitempty"`
	FailureMessage string `gorm:"size:500" json:"failure_message,omitempty"`

	// Timestamps
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ProcessedAt *time.Time     `json:"processed_at,omitempty"`

	// Relationships
	Refunds []Refund `gorm:"foreignKey:PaymentID" json:"refunds,omitempty"`
}

// Enums
type PaymentStatus string

const (
	PaymentStatusPending    PaymentStatus = "pending"
	PaymentStatusProcessing PaymentStatus = "processing"
	PaymentStatusCompleted  PaymentStatus = "completed"
	PaymentStatusFailed     PaymentStatus = "failed"
	PaymentStatusCancelled  PaymentStatus = "cancelled"
	PaymentStatusRefunded   PaymentStatus = "refunded"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodDebitCard    PaymentMethod = "debit_card"
	PaymentMethodPayPal       PaymentMethod = "paypal"
	PaymentMethodStripe       PaymentMethod = "stripe"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodCrypto       PaymentMethod = "crypto"
)
