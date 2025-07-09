package domain

import (
	"time"
)

type Transaction struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Amount      float64   `gorm:"type:numeric(10,2);not null" json:"amount" binding:"required"`
    Currency    string    `gorm:"type:varchar(10);not null" json:"currency" binding:"required"`
    Email       string    `gorm:"type:varchar(255);not null" json:"customer_email" binding:"required,email"`
    TxnRef      string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"txn_ref" binding:"required"`
    Status      string    `gorm:"type:varchar(100);default:'pending'" json:"status,omitempty"`
    Description string    `gorm:"type:text" json:"description,omitempty"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

