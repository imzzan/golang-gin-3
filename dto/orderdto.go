package dto

import "time"

type OrderDto struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}

type OrderResponse struct {
	Id        string          `json:"id"`
	User      UserResponse    `json:"user"`
	Product   ProductResponse `json:"product"`
	IsPaid    bool            `json:"is_paid"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
