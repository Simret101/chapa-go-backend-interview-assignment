package domain

import (
	"time"
)

type WebhookLog struct {
	ID         uint      `gorm:"primaryKey"`
	Event      string    `gorm:"type:varchar(255)"`
	Payload    string    `gorm:"type:jsonb"`
	StatusCode int       `gorm:"default:200"`
	ReceivedAt time.Time `gorm:"autoCreateTime"`
}
