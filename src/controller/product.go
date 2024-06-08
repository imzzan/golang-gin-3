package controller

import (
	"fmt"
	"golang-gin3/dto"
	"golang-gin3/errorhandler"
	"golang-gin3/helper"
	"golang-gin3/src/service"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *productController {
	return &productController{productService}
}

func (c *productController) Create(ctx *gin.Context) {
	payload := dto.ProductDto{}

	err := ctx.ShouldBind(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	ext := filepath.Ext(payload.Image.Filename)
	newFileName := uuid.New().String() + ext

	// Save the image
	dst := filepath.Join("public", filepath.Base(newFileName))
	ctx.SaveUploadedFile(payload.Image, dst)

	userId, _ := ctx.Get("user_id")
	payload.UserId = userId.(string)
	payload.Image.Filename = fmt.Sprintf("%s/public/%s", ctx.Request.Host, newFileName)

	product, err := c.productService.Create(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success Created Product",
		Data:       product,
	})

	ctx.JSON(http.StatusCreated, response)
}

func (c *productController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.FindAll()
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusAccepted,
		Message:    "Succeess",
		Data:       products,
	})

	ctx.JSON(http.StatusCreated, response)
}

func (c *productController) FindById(ctx *gin.Context) {
	id, _ := ctx.Get("product_id")
	productId := id.(string)

	fmt.Println(id)
	product, err := c.productService.FindById(productId)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success Product By Id",
		Data:       product,
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *productController) FindByUserId(ctx *gin.Context) {
	id, _ := ctx.Get("user_id")
	userId := id.(string)
	products, err := c.productService.FindByUserId(userId)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       products,
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *productController) Update(ctx *gin.Context) {
	payload := dto.UpdateProductDto{}
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		errorhandler.HandlerError(ctx, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}
	id, _ := ctx.Get("product_id")
	productId := id.(string)

	product, err := c.productService.Update(&payload, productId)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusCreated,
		Message:    "Success updated",
		Data:       product,
	})

	ctx.JSON(http.StatusCreated, response)
}

func (c *productController) Delete(ctx *gin.Context) {
	id, _ := ctx.Get("product_id")
	productId := id.(string)

	err := c.productService.Delete(productId)
	if err != nil {
		errorhandler.HandlerError(ctx, err)
		return
	}

	response := helper.Response(dto.ResponsePrams{
		StatusCode: http.StatusOK,
		Message:    "Product success deleted",
	})

	ctx.JSON(http.StatusOK, response)
}
