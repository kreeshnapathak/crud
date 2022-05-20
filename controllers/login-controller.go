package controllers

import (
	// "crud/dto"
	"crud/models"
	"crud/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	jWtService services.JWTService
}

func LoginHandler(jWtService services.JWTService) LoginController {
	return &loginController{
		jWtService: jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "Inproper data"
	}

	fmt.Println(user.Email, user.Password, "from login controller")
	isUserAuthenticated := services.LoginUser(user.Email, user.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(user.Email, true)

	}
	return "Invalid credentials"
}
