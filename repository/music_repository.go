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
		music, err := scanToMusic(rows)

		helpers.PanicIfError(err)

		musics = append(musics, *music)
	}

	return musics
}

func (mr *MusicRepository) GetMusicsByPublished(ctx context.Context, tx *sql.Tx, isPublished bool) []models.Music {

	rawSql := "select id,title,artist,is_published from musics where is_published = ?"

	rows, err := tx.QueryContext(ctx, rawSql, isPublished)

	helpers.PanicIfError(err)

	defer rows.Close()

	var musics []models.Music

	for rows.Next() {

		music, err := scanToMusic(rows)

		helpers.PanicIfError(err)

		musics = append(musics, *music)
	}

	return musics
}

func (mr *MusicRepository) GetMusicById(ctx context.Context, tx *sql.Tx, id string) (models.Music, error) {

	rawSql := "select id,title,artist,is_published from musics where id = ?"

	rows, err := tx.QueryContext(ctx, rawSql, id)

	helpers.PanicIfError(err)

	defer rows.Close()

	if rows.Next() {

		music, err := scanToMusic(rows)

		helpers.PanicIfError(err)

		return *music, nil

	} else {
		return models.Music{}, errors.New("music not found")
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

func (mr *MusicRepository) UpdateMusic(ctx context.Context, tx *sql.Tx, music models.Music) models.Music {

	rawSql := "UPDATE musics SET title = ?,artist = ? ,is_published = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, rawSql, music.Title, music.Artist, music.IsPublished, music.ID)

	helpers.PanicIfError(err)

	return music
}

func scanToMusic(rows *sql.Rows) (*models.Music, error) {

	music := models.Music{}

	err := rows.Scan(&music.ID, &music.Title, &music.Artist, &music.IsPublished)

	return &music, err

}
