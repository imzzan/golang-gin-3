package errorhandler

import (
	"golang-gin3/dto"
	"golang-gin3/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerError(ctx *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case *ForBiddenError:
		statusCode = http.StatusForbidden
	case *ConflictError:
		statusCode = http.StatusConflict
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	ctx.JSON(statusCode, response)
}
