package services

type LoginService interface {
	Login(email string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "username",
		authorizedPassword: "password",
	}
}
func (service *loginService) Login(email string, password string) bool {
	return service.authorizedUsername == email && service.authorizedPassword == password
}

// func (service *loginService) Login(email string, password string) bool {
// 	return info.email == email && info.password == password
// 	var user models.User
// 	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
// 		return false
// 	} else if email == service.authorizedUsername && password == service.authorizedPassword {
// 		return true
// 	} else {
// 		return false
// 	}
// }
