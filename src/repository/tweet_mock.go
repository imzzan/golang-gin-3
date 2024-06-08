package repository

import (
	"errors"
	"golang-gin3/schema"

	"github.com/stretchr/testify/mock"
)

type TweetRepoMock struct {
	Mock mock.Mock
}

func (r *TweetRepoMock) Create(payload *schema.Tweet) error {
	arguments := r.Mock.Called(&payload)

	if arguments.Get(0) == nil {
		return errors.New("Error")
	}

	_ = arguments.Get(0).(schema.Tweet)
	return nil
}

func (r *TweetRepoMock) FindAll() ([]*schema.Tweet, error) {
	var tweet []*schema.Tweet
	arguments := r.Mock.Called(&tweet)
	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	}
	_ = arguments.Get(0).(schema.Tweet)
	return tweet, nil

}

func (repository *TweetRepoMock) FindById(id string) (*schema.Tweet, error) {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	}

	tweet := arguments.Get(0).(schema.Tweet)
	return &tweet, nil
}

func (r *TweetRepoMock) Update(payload *schema.Tweet) error {
	arguments := r.Mock.Called(&payload)

	if arguments.Get(0) == nil {
		return errors.New("Error")
	}

	_ = arguments.Get(0).(schema.Tweet)
	return nil
}

func (r *TweetRepoMock) Delete(id string) error {
	arguments := r.Mock.Called(id)

	if arguments.Get(0) == nil {
		return errors.New("Error")
	}

	_ = arguments.Get(0).(schema.Tweet)
	return nil
}
