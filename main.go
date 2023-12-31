package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/radhitka/go-music/config"
	"github.com/radhitka/go-music/controllers"
	"github.com/radhitka/go-music/repository"
	"github.com/radhitka/go-music/response"
	"github.com/radhitka/go-music/services"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.LoadDatabase()

	router := gin.Default()
	router.HandleMethodNotAllowed = true

	musicRepository := repository.NewMusicRepository()
	musicService := services.NewMusicService(*musicRepository, db)
	musicController := controllers.NewMusicController(*musicService)

	router.GET("/musics", musicController.GetMusics)
	router.POST("/musics", musicController.AddMusic)
	router.GET("/musics/:id", musicController.GetMusicById)
	router.PUT("/musics/:id", musicController.UpdateMusic)
	router.DELETE("/musics/:id", musicController.DeleteMusic)
	router.NoRoute(handleNoRoute)
	router.NoMethod(handleNoMethod)

	fmt.Println("Server running.....")

	router.Run("localhost:8080")

}

func handleNoRoute(c *gin.Context) {
	notFoundResponse := response.NewResponseData().NotFound().WithMessage("Route Not found!")

	c.IndentedJSON(notFoundResponse.Code, notFoundResponse)
}

func handleNoMethod(c *gin.Context) {
	notFoundResponse := response.NewResponseData().MethodNotAllowed().WithMessage("Method Not Allowed!")

	c.IndentedJSON(notFoundResponse.Code, notFoundResponse)
}
