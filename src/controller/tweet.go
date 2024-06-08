package controller

import (
	"golang-gin3/dto"
	"golang-gin3/errorhandler"
	"golang-gin3/helper"
	"golang-gin3/src/service"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweetService service.TweetService
}

func NewTweetController(tweetService service.TweetService) *tweetController {
	return &tweetController{tweetService}
}

func (c *tweetController) Create(ctx *gin.Context) {
	payload := dto.TweetCreateDto{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	userId, _ := ctx.Get("user_id")
	payload.UserId = userId.(string)

	tweet, err := c.tweetService.Create(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusCreated,
		Message:    "Create tweet success",
		Data:       tweet,
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *tweetController) FindAll(ctx *gin.Context) {
	tweets, err := c.tweetService.FindAll()

	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       tweets,
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *tweetController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	tweet, err := c.tweetService.FindById(id)
	if err != nil {
		log.Info("Tweet not found")
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       *tweet,
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *tweetController) Update(ctx *gin.Context) {
	var payload dto.UpdateTweet

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	id := ctx.Param("id")

	tweet, err := c.tweetService.Update(&payload, id)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       tweet,
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.tweetService.Delete(id)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Delete Successfully",
	})

	ctx.JSON(http.StatusOK, response)
}
