package main

import (
	"fmt"
	"net/http"

	"github.com/radhitka/go-music/models"
	"github.com/radhitka/go-music/response"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/musics", getMusics)
	router.GET("/musics/:id", getMusicById)

	fmt.Println("Server running.....")

	router.Run("localhost:8080")

}

func getMusics(c *gin.Context) {

	mr := models.Musics

	newReponse := response.NewSuccessResponse(response.ToMusicsReponse(mr))

	c.IndentedJSON(http.StatusOK, newReponse)
}

func getMusicById(c *gin.Context) {

	id := c.Param("id")

	mr := models.Musics

	for _, m := range mr {
		if m.ID == id {

			successResponse := response.NewSuccessResponse(response.ToMusicReponse(m))
			c.IndentedJSON(successResponse.Code, successResponse)
			return
		}
	}

	notFoundResponse := response.NewNotFoundResponse()

	c.IndentedJSON(notFoundResponse.Code, notFoundResponse)

}
