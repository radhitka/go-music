package request

type MusicRequest struct {
	Name string `form:"name" binding:"required"`
}
