package response

import "net/http"

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponseData() *ResponseData {
	return &ResponseData{}
}

func (res *ResponseData) WithCode(code int) *ResponseData {
	res.Code = code
	return res
}

func (res *ResponseData) WithData(data any) *ResponseData {
	res.Data = data
	return res
}

func (res *ResponseData) RemoveData() *ResponseData {
	res.Data = nil
	return res
}

func (res *ResponseData) WithMessage(statusMessage string) *ResponseData {
	res.Message = statusMessage
	return res
}

func (res *ResponseData) Success() *ResponseData {
	return res.WithCode(http.StatusOK).WithMessage(http.StatusText(http.StatusOK))
}

func (res *ResponseData) Error() *ResponseData {
	return res.WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
}

func (res *ResponseData) SuccessCreated() *ResponseData {
	return res.WithCode(http.StatusCreated).WithMessage(http.StatusText(http.StatusCreated))
}

func (res *ResponseData) NotFound() *ResponseData {
	return res.WithCode(http.StatusNotFound).WithMessage(http.StatusText(http.StatusNotFound))
}

func (res *ResponseData) Forbidden() *ResponseData {
	return res.WithCode(http.StatusForbidden).WithMessage(http.StatusText(http.StatusForbidden))
}

func (res *ResponseData) MethodNotAllowed() *ResponseData {
	return res.WithCode(http.StatusMethodNotAllowed).WithMessage(http.StatusText(http.StatusMethodNotAllowed))
}

func (res *ResponseData) UnprocessableEntity() *ResponseData {
	return res.WithCode(http.StatusUnprocessableEntity)
}
