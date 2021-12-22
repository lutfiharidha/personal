package models

import (
	"time"
)

type Content struct {
	ID                 string    `json:"id" gorm:"uniqueIndex;not null"`
	ContentTitle       string    `json:"content_title" gorm:"not null"`
	ContentImage       string    `json:"content_image"`
	ContentDescription string    `json:"content_description" gorm:"not null"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
