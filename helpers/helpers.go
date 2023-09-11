package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/radhitka/go-music/response"
)

func HandleValidationErr(c *gin.Context, err error) {

	for _, fieldErr := range err.(validator.ValidationErrors) {
		res := response.NewErrorResponse(http.StatusBadRequest, fieldErr.Error())

		c.IndentedJSON(res.Code, res)
		return
	}

}
