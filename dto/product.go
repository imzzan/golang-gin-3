package dto

import (
	"mime/multipart"
	"time"
)

type ProductDto struct {
	Name   string                `form:"name"`
	Price  int                   `form:"price"`
	Image  *multipart.FileHeader `form:"image"`
	UserId string                `json:"user_id"`
}

type UpdateProductDto struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UserProduct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateProductResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Image     string `json:"image"`
	User      UserProduct
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
