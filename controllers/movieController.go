package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	model "github.com/BhandariG29/movie-recommendation-system/models"
	database "github.com/BhandariG29/movie-recommendation-system/database"
)

var validate = validator.New()

func GetMovies() gin.HandlerFunc {
	return func (c *gin.Context){
		var movies []model.Movie

		database.DB.
			Preload("Director").
			Preload("Genres").
			Preload("Actors").
			Find(&movies)

		c.JSON(http.StatusOK, movies)
	}
}

func CreateMovie() gin.HandlerFunc {
	return func (c *gin.Context){
		var movie model.Movie

		if err := c.BindJSON(&movie); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		validationErr := validate.Struct(movie)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return 
		}

		result := database.DB.Create(&movie)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return 
		}

		c.JSON(http.StatusOK, movie)
	}
}