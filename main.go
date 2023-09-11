package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/radhitka/go-music/config"
	"github.com/radhitka/go-music/controllers"
	"github.com/radhitka/go-music/response"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.LoadDatabase()

	router := gin.New()

	musicController := controllers.NewMusicController(db)

	router.GET("/musics", musicController.GetMusics)
	router.POST("/musics", musicController.AddMusic)
	router.GET("/musics/:id", musicController.GetMusicById)
	router.NoRoute(handleNoRoute)

	fmt.Println("Server running.....")

	router.Run("localhost:8080")

}

func handleNoRoute(c *gin.Context) {
	notFoundResponse := response.NewNotFoundResponse("Route not found!")

	c.IndentedJSON(notFoundResponse.Code, notFoundResponse)
}
