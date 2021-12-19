package repositories

import (
	"fmt"
	"personal/app/configs"
	"personal/app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterAuth(au models.Auth) models.Auth
	LoginCheck(username string, password string) (string, error)
}

type authConnection struct {
	connection *gorm.DB
}

func NewAuthRepository(dbConn *gorm.DB) AuthRepository {
	return &authConnection{
		connection: dbConn,
	}
}

func (db *authConnection) RegisterAuth(auth models.Auth) models.Auth {
	db.connection.Create(&auth)
	db.connection.Find(&auth)
	return auth
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (db *authConnection) LoginCheck(username string, password string) (string, error) {

	var err error

	u := models.Auth{}

	err = db.connection.Model(models.Auth{}).Where("email = ?", username).Or("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := configs.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}
	fmt.Println(token)
	return token, nil

}
