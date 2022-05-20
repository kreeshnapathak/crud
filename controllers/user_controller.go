package controllers

import (
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create task
	user := models.User{ID: input.ID, Name: input.Name, Email: input.Email, Password: input.Password, IsVerified: input.IsVerified}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
