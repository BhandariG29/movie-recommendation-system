package main

import (
	"log"
	"github.com/gin-gonic/gin"

	database "github.com/BhandariG29/movie-recommendation-system/database"
	model "github.com/BhandariG29/movie-recommendation-system/models"
	// route "github.com/BhandariG29/movie-recommendation-system/routes"
)

func main(){
	database.ConnectionDB()
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance")
	}

	defer sqlDB.Close()

	database.DB.AutoMigrate(
		&model.User{},
		&model.Movie{},
		&model.Actor{},
		&model.Director{},
		&model.Genre{},
		&model.MovieActor{},
		&model.MovieGenre{},
		&model.WatchHistory{},
	)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Movie Recommendation API Running",
		})
	})
	// router.GET("/users", func(c *gin.Context){

	// })
	// route.UserRoute(router)
	router.Run(":8080")
}