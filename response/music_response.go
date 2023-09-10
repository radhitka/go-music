package response

import (
	"github.com/radhitka/go-music/models"
)

type MusicResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	IsPublished bool   `json:"is_published"`
}

func ToMusicReponse(m models.Music) MusicResponse {

	return MusicResponse{
		ID:          m.ID,
		Title:       m.Title,
		Artist:      m.Artist,
		IsPublished: true,
	}
}

func ToMusicsReponse(ms []models.Music) []MusicResponse {

	var mr []MusicResponse

	for _, m := range ms {
		mr = append(mr, ToMusicReponse(m))
	}

	return mr
}
