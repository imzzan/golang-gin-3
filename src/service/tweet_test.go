package service

import (
	"golang-gin3/schema"
	"golang-gin3/src/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var tweetRepository = &repository.TweetRepoMock{Mock: mock.Mock{}}
var tweetSevice = tweetService{tweetRepo: tweetRepository}

func TestTweetFindByIdServiceNotFound(t *testing.T) {
	tweetRepository.Mock.On("FindById", "1").Return(nil)
	tweet, err := tweetSevice.FindById("1")
	assert.NotNil(t, err)
	assert.Nil(t, tweet)
}

func TestTweetFindIdServiceFound(t *testing.T) {
	tweet := schema.Tweet{
		Id:        "1",
		Title:     "This is the title",
		Caption:   "This is the caption",
		Image:     "This is the image",
		UserId:    "userId",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tweetRepository.Mock.On("FindById", "2").Return(tweet)
	result, err := tweetSevice.FindById("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
