package services

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/radhitka/go-music/helpers"
	"github.com/radhitka/go-music/models"
	"github.com/radhitka/go-music/repository"
	"github.com/radhitka/go-music/request"
	"github.com/radhitka/go-music/response"
)

type MusicService struct {
	MusicRepository repository.MusicRepository
	DB              *sql.DB
}

func NewMusicService(mr repository.MusicRepository, db *sql.DB) *MusicService {
	return &MusicService{
		MusicRepository: mr,
		DB:              db,
	}
}

func (ms *MusicService) GetMusics(c *gin.Context) ([]response.MusicResponse, error) {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	musics, err := ms.MusicRepository.GetMusics(c.Request.Context(), tx)

	return response.ToMusicsResponse(musics), err
}

func (ms *MusicService) GetMusicsByFiltered(c *gin.Context) ([]response.MusicResponse, error) {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	musics, err := ms.MusicRepository.GetMusics(c.Request.Context(), tx)

	if c.Query("published") != "" {
		isPublished, _ := strconv.ParseBool(c.Query("published"))

		musics, err = ms.MusicRepository.GetMusicsByPublished(c.Request.Context(), tx, isPublished)
	}

	return response.ToMusicsResponse(musics), err
}

func (ms *MusicService) GetMusicById(c *gin.Context, id string) (response.MusicResponse, error) {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	music, err := ms.MusicRepository.GetMusicById(c.Request.Context(), tx, id)

	return response.ToMusicResponse(music), err
}

func (ms *MusicService) AddMusic(c *gin.Context, req request.MusicRequest) (response.MusicResponse, error) {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	music := models.Music{
		Title:       req.Title,
		Artist:      req.Artist,
		IsPublished: req.IsPublised,
	}

	music, err = ms.MusicRepository.AddMusic(c.Request.Context(), tx, music)

	return response.ToMusicResponse(music), err
}

func (ms *MusicService) UpdateMusic(c *gin.Context, req request.MusicRequest, id string) (response.MusicResponse, error) {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	music, err := ms.MusicRepository.GetMusicById(c.Request.Context(), tx, id)

	if err != nil {
		return response.ToMusicResponse(music), err
	}

	newId, _ := strconv.Atoi(id)

	music.ID = newId
	music.Title = req.Title
	music.Artist = req.Artist
	music.IsPublished = req.IsPublised

	music, err = ms.MusicRepository.UpdateMusic(c.Request.Context(), tx, music)

	return response.ToMusicResponse(music), err
}

func (ms *MusicService) DeleteMusic(c *gin.Context, id string) error {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	music, err := ms.MusicRepository.GetMusicById(c.Request.Context(), tx, id)

	if err != nil {
		return err
	}

	newId, _ := strconv.Atoi(id)

	music.ID = newId

	return ms.MusicRepository.DeleteMusic(c.Request.Context(), tx, music)
}
