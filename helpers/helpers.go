package helpers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/radhitka/go-music/response"
)

func HandleValidationErr(c *gin.Context, err error) {
	for _, fieldErr := range err.(validator.ValidationErrors) {
		res := response.NewResponseData().UnprocessableEntity().WithMessage(fieldErr.Error())

		c.IndentedJSON(res.Code, res)
		return
	}
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
