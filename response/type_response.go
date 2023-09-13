package response

import "net/http"

type responseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponseData() *responseData {
	return &responseData{}
}

func (res *responseData) WithCode(code int) *responseData {
	res.Code = code
	return res
}

func (res *responseData) WithData(data any) *responseData {
	res.Data = data
	return res
}

func (res *responseData) WithMessage(statusMessage string) *responseData {
	res.Message = statusMessage
	return res
}

func (res *responseData) Success() *responseData {
	return res.WithCode(http.StatusOK).WithMessage(http.StatusText(http.StatusOK))
}

func (res *responseData) SuccessCreated() *responseData {
	return res.WithCode(http.StatusCreated).WithMessage(http.StatusText(http.StatusCreated))
}

func (res *responseData) NotFound() *responseData {
	return res.WithCode(http.StatusNotFound).WithMessage(http.StatusText(http.StatusNotFound))
}

func (res *responseData) Forbidden() *responseData {
	return res.WithCode(http.StatusForbidden).WithMessage(http.StatusText(http.StatusForbidden))
}

func (res *responseData) MethodNotAllowed() *responseData {
	return res.WithCode(http.StatusMethodNotAllowed).WithMessage(http.StatusText(http.StatusMethodNotAllowed))
}

func (res *responseData) UnprocessableEntity() *responseData {
	return res.WithCode(http.StatusUnprocessableEntity)
}
