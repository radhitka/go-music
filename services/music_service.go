package services

import (
	"context"
	"database/sql"

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
