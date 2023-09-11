package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/radhitka/go-music/helpers"
	"github.com/radhitka/go-music/models"
	"github.com/radhitka/go-music/request"
	"github.com/radhitka/go-music/response"
)

type MusicController struct {
	DB *sql.DB
}

func NewMusicController(db *sql.DB) *MusicController {
	return &MusicController{
		DB: db,
	}
}

func (ms *MusicController) GetMusics(c *gin.Context) {

	rows, err := ms.DB.Query("select id,title,artist,is_published from musics")

	if err != nil {
		log.Fatal(err)

	}

	defer rows.Close()

	var musics []models.Music

	for rows.Next() {

		music := models.Music{}

		err := rows.Scan(&music.ID, &music.Title, &music.Artist, &music.IsPublished)

		if err != nil {
			log.Fatal(err)
		}

		musics = append(musics, music)

	}

	newReponse := response.NewSuccessResponse(response.ToMusicsReponse(musics))

	c.IndentedJSON(newReponse.Code, newReponse)
}

func (ms *MusicController) GetMusicById(c *gin.Context) {
	id := c.Param("id")

	music := models.Music{}

	rsql := "select id,title,artist,is_published from musics where id = ?"

	err := ms.DB.QueryRow(rsql, id).Scan(&music.ID, &music.Title, &music.Artist, &music.IsPublished)

	switch err {

	case sql.ErrNoRows:
		notFoundResponse := response.NewNotFoundResponse("Music not found")

		c.IndentedJSON(notFoundResponse.Code, notFoundResponse)
		return
	case nil:
		successResponse := response.NewSuccessResponse(response.ToMusicReponse(music))

		c.IndentedJSON(successResponse.Code, successResponse)
		return
	default:
		log.Fatal(err)
	}
}

func (ms *MusicController) AddMusic(c *gin.Context) {

	var musicRequest request.MusicRequest

	err := c.ShouldBind(&musicRequest)

	if err != nil {

		helpers.HandleValidationErr(c, err)
		return
	}

	c.IndentedJSON(200, c.PostForm("name"))
}
