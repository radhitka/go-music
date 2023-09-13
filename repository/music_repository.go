package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/radhitka/go-music/helpers"
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

	helpers.PanicIfError(err)

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

	helpers.PanicIfError(err)

	defer rows.Close()

	music := models.Music{}
	if rows.Next() {
		err := rows.Scan(&music.ID, &music.Title, &music.Artist, &music.IsPublished)

		helpers.PanicIfError(err)

		return music, nil

	} else {
		return music, errors.New("music not found")
	}
}

func (mr *MusicRepository) AddMusic(ctx context.Context, tx *sql.Tx, music models.Music) models.Music {

	rawSql := "insert into musics(title,artist,is_published) values (?, ?, ?)"

	result, err := tx.ExecContext(ctx, rawSql, music.Title, music.Artist, music.IsPublished)

	helpers.PanicIfError(err)

	newId, err := result.LastInsertId()

	helpers.PanicIfError(err)

	music.ID = int(newId)

	return music
}
