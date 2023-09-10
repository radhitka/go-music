package response

import "net/http"

type ReponseData struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func NewSuccessResponse(d any) *ReponseData {

	return &ReponseData{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   d,
	}

}

func NewNotFoundResponse() *ReponseData {

	return &ReponseData{
		Code:   http.StatusNotFound,
		Status: "NotFound",
	}

}
