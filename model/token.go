package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Token type
type Token struct {
	ID              int `gorm:"primary_key" json:"id"`
	UserID          int
	UUID            uuid.UUID `gorm:"type:uuid;type:varchar(100);"`
	Token           string    `gorm:"type:text;"`
	TransmissionKey string    `gorm:"type:text;"`
	ExpiryTime      time.Time
}
