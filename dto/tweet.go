package dto

import "time"

type UserTweet struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TweetCreateDto struct {
	Title   string `json:"title"`
	Caption string `json:"caption"`
	Image   string `json:"image"`
	UserId  string `json:"user_id"`
}

type TweetResponse struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Image     string `json:"image"`
	User      UserTweet
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateTweet struct {
	Title   string `json:"title"`
	Caption string `json:"caption"`
	Image   string `json:"image"`
}
