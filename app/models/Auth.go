package models

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	ID        string         `gorm:"uniqueIndex" json:"id"`
	Email     string         `json:"email" gorm:"unique:not null"`
	Password  string         `gorm:"not null" json:"password"`
	Username  string         `json:"username" gorm:"unique:not null"`
	Name      string         `json:"name" gorm:"not null"`
	Image     string         `json:"image"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"  json:"deleted_at"`
}

func (auth *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	username := html.EscapeString(strings.TrimSpace(auth.Username))
	tx.Statement.SetColumn("ID", uuid)
	tx.Statement.SetColumn("Password", hashedPassword)
	tx.Statement.SetColumn("Username", username)
	return
}
