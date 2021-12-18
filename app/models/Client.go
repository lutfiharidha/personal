package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID          string         `gorm:"uniqueIndex" json:"id"`
	ClientName  string         `json:"client_name" gorm:"not null"`
	ClientImage string         `gorm:"index;not null" json:"client_image"`
	ClientTitle string         `json:"client_title" gorm:"not null"`
	Description string         `json:"description"`
	Status      uint8          `json:"status" gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"  json:"deleted_at"`
}

func (product *Client) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
