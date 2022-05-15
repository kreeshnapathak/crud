package controllers

import (
	"crud/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//getting all the movies list
func GetMovies(c *gin.Context) {
	fmt.Println("GetMovies----------")
	var movies []models.Movie
	models.DB.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// Find a movie by id
func GetMovie(c *gin.Context) { // Get model if exist
	var movie models.Movie
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

//creating a new movie
func CreateMovie(c *gin.Context) {
	// Validate input
	var input models.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create task
	movie := models.Movie{ID: input.ID, MovieName: input.MovieName}
	models.DB.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Update a movie
func UpdateMovie(c *gin.Context) {
	var movie models.Movie
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&movie).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Delete a movie
func DeleteMovie(c *gin.Context) {
	var movie models.Movie
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}
