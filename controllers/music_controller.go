package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/radhitka/go-music/request"
	"github.com/radhitka/go-music/response"
	"github.com/radhitka/go-music/services"
)

type MusicController struct {
	MusicService services.MusicService
}

func NewMusicController(ms services.MusicService) *MusicController {
	return &MusicController{
		MusicService: ms,
	}
}

func (mc *MusicController) GetMusics(c *gin.Context) {

	musics := mc.MusicService.GetMusics(c.Request.Context())

	if len(c.Request.URL.Query()) > 0 {
		musics = mc.MusicService.GetMusicsByFiltered(c.Request.Context(), c)
	}

	res := response.NewResponseData().Success().WithData(musics)

	c.IndentedJSON(res.Code, res)

}

func (mc *MusicController) GetMusicById(c *gin.Context) {
	id := c.Param("id")

	music, empty := mc.MusicService.GetMusicById(c.Request.Context(), id)

	res := response.NewResponseData().Success().WithData(music)

	if empty {
		res = response.NewResponseData().NotFound().WithMessage("Music Not Found!")
	}

	c.IndentedJSON(res.Code, res)
}

func (mc *MusicController) AddMusic(c *gin.Context) {
	var musicRequest request.MusicRequest

	c.Bind(&musicRequest)

	music := mc.MusicService.AddMusic(c.Request.Context(), musicRequest)

	res := response.NewResponseData().SuccessCreated().WithData(music)

	c.IndentedJSON(res.Code, res)
}

func (mc *MusicController) UpdateMusic(c *gin.Context) {
	id := c.Param("id")

	var musicRequest request.MusicRequest

	c.Bind(&musicRequest)

	music := mc.MusicService.UpdateMusic(c.Request.Context(), musicRequest, id)

	res := response.NewResponseData().SuccessCreated().WithData(music)

	c.IndentedJSON(res.Code, res)

}
