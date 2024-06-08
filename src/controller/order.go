package controller

import (
	"golang-gin3/dto"
	"golang-gin3/errorhandler"
	"golang-gin3/helper"
	"golang-gin3/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *orderController {
	return &orderController{orderService: orderService}
}

func (c *orderController) Create(ctx *gin.Context) {
	paylaod := dto.OrderDto{}

	err := ctx.ShouldBindJSON(&paylaod)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.BadRequestError{Message: "Bad Request"})
		return
	}

	userId, _ := ctx.Get("user_id")
	paylaod.UserId = userId.(string)

	err = c.orderService.Create(paylaod)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.BadRequestError{Message: "Bad Request"})
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success",
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) FindAll(ctx *gin.Context) {
	orders, err := c.orderService.FindAll()
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       orders,
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := c.orderService.FindById(id)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       order,
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	paylaod := dto.OrderDto{}

	err := ctx.ShouldBindJSON(&paylaod)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.BadRequestError{Message: "Bad Request"})
		return
	}

	userId, _ := ctx.Get("user_id")
	paylaod.UserId = userId.(string)

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Succes",
	})

	err = c.orderService.Update(id, paylaod)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.orderService.Delete(id)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success",
	})

	ctx.JSON(http.StatusOK, response)
}
