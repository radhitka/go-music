package services

import (
	"context"
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

func (ms *MusicService) GetMusics(ctx context.Context) []response.MusicResponse {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	musics := ms.MusicRepository.GetMusics(ctx, tx)

	return response.ToMusicsResponse(musics)
}

func (ms *MusicService) GetMusicsByFiltered(ctx context.Context, c *gin.Context) []response.MusicResponse {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	if c.Query("published") != "" {
		isPublished, _ := strconv.ParseBool(c.Query("published"))
		musics := ms.MusicRepository.GetMusicsByPublished(ctx, tx, isPublished)

		return response.ToMusicsResponse(musics)
	}

	musics := ms.MusicRepository.GetMusics(ctx, tx)

	return response.ToMusicsResponse(musics)

}

func (ms *MusicService) GetMusicById(ctx context.Context, id string) (response.MusicResponse, bool) {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	music, err := ms.MusicRepository.GetMusicById(ctx, tx, id)

	if err != nil {
		return response.ToMusicResponse(music), true
	}

	return response.ToMusicResponse(music), false
}

func (ms *MusicService) AddMusic(ctx context.Context, req request.MusicRequest) response.MusicResponse {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	music := models.Music{
		Title:       req.Title,
		Artist:      req.Artist,
		IsPublished: req.IsPublised,
	}

	music = ms.MusicRepository.AddMusic(ctx, tx, music)

	return response.ToMusicResponse(music)
}

func (ms *MusicService) UpdateMusic(ctx context.Context, req request.MusicRequest, id string) response.MusicResponse {

	tx, err := ms.DB.Begin()

	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	music, err := ms.MusicRepository.GetMusicById(ctx, tx, id)

	helpers.PanicIfError(err)

	newId, _ := strconv.Atoi(id)

	music.ID = newId
	music.Title = req.Title
	music.Artist = req.Artist
	music.IsPublished = req.IsPublised

	music = ms.MusicRepository.UpdateMusic(ctx, tx, music)

	return response.ToMusicResponse(music)

}
