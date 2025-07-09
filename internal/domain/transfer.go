package domain

import (
	"time"
)

type Transfer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TransferRef string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"transfer_ref" binding:"required"`
	Account     string    `gorm:"type:varchar(255);not null" json:"account" binding:"required"`
	Amount      float64   `gorm:"type:numeric(10,2);not null" json:"amount" binding:"required"`
	BankCode    string    `gorm:"type:varchar(100);not null" json:"bank_code" binding:"required"`
	Status      string    `gorm:"type:varchar(100);default:'initiated'" json:"status,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
