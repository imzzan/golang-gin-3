package middleware

import (
	"golang-gin3/config"
	"golang-gin3/errorhandler"
	"golang-gin3/helper"
	"golang-gin3/src/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			errorhandler.HandlerError(ctx, &errorhandler.UnauthorizedError{Message: "Unauthorized"})
			return
		}

		jwtHeader := strings.Split(header, " ")
		token := jwtHeader[1]

		user, err := helper.ValidateToken(token)
		if err != nil {
			errorhandler.HandlerError(ctx, &errorhandler.UnauthorizedError{Message: "UnAuthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", user.Id)
		ctx.Next()
	}
}

func JustMe() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, _ := ctx.Get("user_id")
		userId := id.(string)

		productId := ctx.Param("id")

		product, err := repository.NewProductRepository(config.Db).FindById(productId)
		if err != nil {
			errorhandler.HandlerError(ctx, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		if product.User.Id != userId {
			errorhandler.HandlerError(ctx, &errorhandler.UnauthorizedError{Message: "Tidak Boleh"})
			return
		}

		ctx.Set("product_id", productId)
		ctx.Next()
	}
}
