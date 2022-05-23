package routes

import (
	"crud/controllers"
	"crud/middleware"
	"crud/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	tokenMaker services.JWTService
}

func SetupRoutes() *gin.Engine {
	var jwtService services.JWTService = services.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(jwtService)

	r := gin.Default()
	fmt.Println("Routes are set up")
	ar := r.Group("/").Use(middleware.AuthMiddleware(jwtService))
	ar.GET("/movies", controllers.GetMovies)
	ar.GET("/movies/:id", controllers.GetMovie)
	ar.POST("/movies", controllers.CreateMovie)
	ar.PATCH("/movies/:id", controllers.UpdateMovie)
	ar.DELETE("movies/:id", controllers.DeleteMovie)
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
