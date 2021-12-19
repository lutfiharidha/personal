package repositories

import (
	"personal/app/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	Profile(uid string) models.Auth
	Logout(uid string) error
}

type profileConnection struct {
	connection *gorm.DB
}

func NewProfileRepository(dbConn *gorm.DB) ProfileRepository {
	return &profileConnection{
		connection: dbConn,
	}
}

func (db *profileConnection) Profile(uid string) models.Auth {
	var user models.Auth
	db.connection.Where("id =?", uid).First(&user)
	return user
}

func (db *profileConnection) Logout(uid string) error {
	var user models.Auth
	db.connection.Where("id = ?", uid).Take(&user).Delete(&user)
	return nil
}
