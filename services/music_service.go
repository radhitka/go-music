package services

import (
	"context"
	"database/sql"

	"github.com/radhitka/go-music/repository"
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

	if err != err {
		panic(err)
	}

	musics := ms.MusicRepository.GetMusics(ctx, tx)

	return response.ToMusicsResponse(musics)
}

func (ms *MusicService) GetMusicById(ctx context.Context, id string) (response.MusicResponse, bool) {

	tx, err := ms.DB.Begin()

	if err != nil {
		panic(err)
	}

	music, err := ms.MusicRepository.GetMusicById(ctx, tx, id)

	if err != nil {
		return response.ToMusicResponse(music), true
	}

	return response.ToMusicResponse(music), false
}

func (ms *MusicService) AddMusic() response.MusicResponse {
	return response.MusicResponse{}

}
