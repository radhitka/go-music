package response

import "net/http"

type reponseData struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type reponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewSuccessResponse(d any) *reponseData {

	return &reponseData{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   d,
	}

}

func NewNotFoundResponse(m string) *reponseWithoutData {

	return &reponseWithoutData{
		Code:    http.StatusNotFound,
		Status:  "NotFound",
		Message: m,
	}

}

func NewErrorResponse(c int, m string) *reponseWithoutData {

	return &reponseWithoutData{
		Code:    c,
		Status:  http.StatusText(c),
		Message: m,
	}

}
