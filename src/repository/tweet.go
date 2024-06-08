package repository

import (
	"golang-gin3/schema"

	"gorm.io/gorm"
)

type TweetInterface interface {
	Create(payload *schema.Tweet) error
	FindAll() ([]*schema.Tweet, error)
	FindById(id string) (*schema.Tweet, error)
	Update(payload *schema.Tweet) error
	Delete(id string) error
}

type tweetRepo struct {
	Db *gorm.DB
}

func NewTweetRepo(db *gorm.DB) *tweetRepo {
	return &tweetRepo{db}
}

func (r *tweetRepo) Create(payload *schema.Tweet) error {
	return r.Db.Create(&payload).Error
}

func (r *tweetRepo) FindAll() ([]*schema.Tweet, error) {
	var tweets []*schema.Tweet
	err := r.Db.Preload("User").Find(&tweets).Error
	return tweets, err
}

func (r *tweetRepo) FindById(id string) (*schema.Tweet, error) {
	var tweets schema.Tweet
	err := r.Db.Preload("User").Find(&tweets, "id=?", id).Error
	return &tweets, err
}

func (r *tweetRepo) Update(payload *schema.Tweet) error {
	err := r.Db.Omit("User").Save(&payload).Error
	return err
}

func (r *tweetRepo) Delete(id string) error {
	var tweet *schema.Tweet
	err := r.Db.Delete(&tweet, "id = ?", id).Error
	return err
}
