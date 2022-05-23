package services

import (
	"crud/helpers"
	"crud/models"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func LoginUser(email string, password string) bool {
	var user models.User
	if fetchError := models.DB.Where("email = ?", email).First(&user).Error; fetchError != nil {
		return false
	}
	// Get hased password from DB
	hashedPassword := []byte(user.Password)
	// Get the password provided in the request.body
	inputpassword := []byte(password)
	err := helpers.PasswordCompare(inputpassword, hashedPassword)
	if err == nil {
		return true
	} else {
		return false
	}

}
