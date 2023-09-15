package repository

import (
	"context"
	"database/sql"

	"github.com/radhitka/go-music/helpers"
	"github.com/radhitka/go-music/models"
)

type MusicRepository struct {
}

func NewMusicRepository() *MusicRepository {
	return &MusicRepository{}
}

func (mr *MusicRepository) GetMusics(ctx context.Context, tx *sql.Tx) ([]models.Music, error) {

	rawSql := "select id,title,artist,is_published from musics"

	rows, err := tx.QueryContext(ctx, rawSql)

	var musics []models.Music

	if err != nil {
		return musics, err
	}

	defer rows.Close()

	for rows.Next() {
		music, err := scanToMusic(rows)

		if err != nil {
			return musics, err
		}

		musics = append(musics, *music)
	}

	return musics, nil
}

func (mr *MusicRepository) GetMusicsByPublished(ctx context.Context, tx *sql.Tx, isPublished bool) ([]models.Music, error) {

	rawSql := "select id,title,artist,is_published from musics where is_published = ?"

	rows, err := tx.QueryContext(ctx, rawSql, isPublished)

	var musics []models.Music

	if err != nil {
		return musics, err
	}

	defer rows.Close()

	for rows.Next() {

		music, err := scanToMusic(rows)

		if err != nil {
			return musics, err
		}

		musics = append(musics, *music)
	}

	return musics, nil
}

func (mr *MusicRepository) GetMusicById(ctx context.Context, tx *sql.Tx, id string) (models.Music, error) {

	rawSql := "select id,title,artist,is_published from musics where id = ?"

	rows, err := tx.QueryContext(ctx, rawSql, id)

	helpers.PanicIfError(err)

	defer rows.Close()

	if rows.Next() {

		music, err := scanToMusic(rows)

		return *music, err
	} else {
		return models.Music{}, sql.ErrNoRows
	}
}

func (mr *MusicRepository) AddMusic(ctx context.Context, tx *sql.Tx, music models.Music) (models.Music, error) {

	rawSql := "insert into musics(title,artist,is_published) values (?, ?, ?)"

	result, err := tx.ExecContext(ctx, rawSql, music.Title, music.Artist, music.IsPublished)

	if err != nil {
		return music, err
	}

	newId, err := result.LastInsertId()

	if err != nil {
		return music, err
	}

	music.ID = int(newId)

	return music, nil
}

func (mr *MusicRepository) UpdateMusic(ctx context.Context, tx *sql.Tx, music models.Music) (models.Music, error) {

	rawSql := "UPDATE musics SET title = ?,artist = ? ,is_published = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, rawSql, music.Title, music.Artist, music.IsPublished, music.ID)

	if err != nil {
		return music, err
	}

	return music, nil
}

func (mr *MusicRepository) DeleteMusic(ctx context.Context, tx *sql.Tx, music models.Music) error {
	rawSql := "delete from musics where id = ?"

	_, err := tx.ExecContext(ctx, rawSql, music.ID)

	return err
}

func scanToMusic(rows *sql.Rows) (*models.Music, error) {

	music := models.Music{}

	err := rows.Scan(&music.ID, &music.Title, &music.Artist, &music.IsPublished)

	return &music, err
}
