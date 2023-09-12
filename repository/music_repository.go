package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/radhitka/go-music/models"
)

type MusicRepository struct {
}

func NewMusicRepository() *MusicRepository {
	return &MusicRepository{}
}

func (mr *MusicRepository) GetMusics(ctx context.Context, tx *sql.Tx) []models.Music {

	rawSql := "select id,title,artist,is_published from musics"

	rows, err := tx.QueryContext(ctx, rawSql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var musics []models.Music

	for rows.Next() {
		music := models.Music{}
		rows.Scan(&music.ID, &music.Title, &music.Artist, &music.IsPublished)

		musics = append(musics, music)
	}

	return musics
}

func (mr *MusicRepository) GetMusicById(ctx context.Context, tx *sql.Tx, id string) (models.Music, error) {

	rawSql := "select id,title,artist,is_published from musics where id = ?"

	rows, err := tx.QueryContext(ctx, rawSql, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	music := models.Music{}
	if rows.Next() {
		err := rows.Scan(&music.ID, &music.Title, &music.Artist, &music.IsPublished)

		if err != nil {
			panic(err)
		}

		return music, nil

	} else {
		return music, errors.New("music not found")
	}
}

func (mr *MusicRepository) AddMusic(ctx context.Context, tx *sql.Tx, music models.Music) models.Music {

	return models.Music{}
}
