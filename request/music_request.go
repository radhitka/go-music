package request

type MusicRequest struct {
	Title      string `form:"title" binding:"required"`
	Artist     string `form:"artist" binding:"required"`
	IsPublised bool   `form:"is_published"`
}
