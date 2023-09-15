package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/radhitka/go-music/helpers"
	"github.com/radhitka/go-music/request"
	"github.com/radhitka/go-music/response"
	"github.com/radhitka/go-music/services"
)

type MusicController struct {
	MusicService services.MusicService
}

const MusicNotFound = "Music Not Found!"

func NewMusicController(ms services.MusicService) *MusicController {
	return &MusicController{
		MusicService: ms,
	}
}

func (mc *MusicController) GetMusics(c *gin.Context) {

	musics, err := mc.MusicService.GetMusics(c)

	if len(c.Request.URL.Query()) > 0 {
		musics, err = mc.MusicService.GetMusicsByFiltered(c)
	}

	res := response.NewResponseData().Success().WithData(musics)

	if err != nil {
		helpers.NewHandleResponseError(err).Handle(res)
	}

	c.IndentedJSON(res.Code, res)

}

func (mc *MusicController) GetMusicById(c *gin.Context) {
	id := c.Param("id")

	music, err := mc.MusicService.GetMusicById(c, id)

	res := response.NewResponseData().Success().WithData(music)

	if err != nil {
		helpers.NewHandleResponseError(err).MessageIfNotFound(MusicNotFound).Handle(res)
	}

	c.IndentedJSON(res.Code, res)
}

func (mc *MusicController) AddMusic(c *gin.Context) {
	var musicRequest request.MusicRequest

	c.Bind(&musicRequest)

	music, err := mc.MusicService.AddMusic(c, musicRequest)

	res := response.NewResponseData().SuccessCreated().WithData(music)

	if err != nil {
		helpers.NewHandleResponseError(err).Handle(res)
	}

	c.IndentedJSON(res.Code, res)
}

func (mc *MusicController) UpdateMusic(c *gin.Context) {
	id := c.Param("id")

	var musicRequest request.MusicRequest

	c.Bind(&musicRequest)

	music, err := mc.MusicService.UpdateMusic(c, musicRequest, id)

	res := response.NewResponseData().SuccessCreated().WithData(music)

	if err != nil {
		helpers.NewHandleResponseError(err).MessageIfNotFound(MusicNotFound).Handle(res)
	}

	c.IndentedJSON(res.Code, res)
}

func (mc *MusicController) DeleteMusic(c *gin.Context) {
	id := c.Param("id")

	err := mc.MusicService.DeleteMusic(c, id)

	res := response.NewResponseData().Success().WithMessage("Success Deleted Music")

	if err != nil {
		helpers.NewHandleResponseError(err).MessageIfNotFound(MusicNotFound).Handle(res)
	}

	c.IndentedJSON(res.Code, res)
}
