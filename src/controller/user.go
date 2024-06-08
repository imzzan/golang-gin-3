package controller

import (
	"golang-gin3/dto"
	"golang-gin3/errorhandler"
	"golang-gin3/helper"
	"golang-gin3/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func UserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (c *userController) Register(ctx *gin.Context) {
	payload := dto.UserDto{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	user, err := c.userService.Create(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusCreated,
		Message:    "User success registered",
		Data:       user,
	})

	ctx.JSON(http.StatusCreated, response)

}

func (c *userController) Login(ctx *gin.Context) {
	payload := dto.LoginDto{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	token, err := c.userService.Login(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "success login",
		Data:       token,
	})

	ctx.JSON(http.StatusOK, response)

}

func (c *userController) GetMe(ctx *gin.Context) {
	id, _ := ctx.Get("user_id")
	userId := id.(string)

	user, err := c.userService.GetMe(userId)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: 200,
		Message:    "Success get me",
		Data:       user,
	})

	ctx.JSON(http.StatusOK, response)
}
