package services

import (
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

	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	} else if email == user.Email && password == user.Password {
		return true
	} else {
		return false
	}
}
