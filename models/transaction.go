package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	Amount          int64     `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	Remarks         string    `json:"remarks,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	BalanceBefore   int64     `json:"balance_before"`
	BalanceAfter    int64     `json:"balance_after"`
}
