package routes

import (
	"crud/controllers"
	"crud/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	var jwtService services.JWTService = services.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(jwtService)

	r := gin.Default()
	fmt.Println("Routes are set up")
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovie)
	r.POST("/movies", controllers.CreateMovie)
	r.PATCH("/movies/:id", controllers.UpdateMovie)
	r.DELETE("movies/:id", controllers.DeleteMovie)
	r.POST("/signup", controllers.CreateUser)
	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	return r
}
