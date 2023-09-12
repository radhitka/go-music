package helpers

import (
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
