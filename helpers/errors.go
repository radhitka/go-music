package helpers

import (
	"database/sql"

	"github.com/radhitka/go-music/response"
)

type HandleResponseError struct {
	err     error
	message string
}

func NewHandleResponseError(e error) *HandleResponseError {
	return &HandleResponseError{
		err:     e,
		message: e.Error(),
	}
}

func (he *HandleResponseError) MessageIfNotFound(text string) *HandleResponseError {
	he.message = text

	return he
}

func (he *HandleResponseError) Handle(res *response.ResponseData) {
	err := he.err

	if err == sql.ErrNoRows {
		res.NotFound().RemoveData().WithMessage(he.message)
	} else if err != sql.ErrNoRows && err != nil {
		res.Error().RemoveData().WithMessage(err.Error())
	}
}
