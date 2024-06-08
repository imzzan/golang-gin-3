package service

import (
	"golang-gin3/dto"
	"golang-gin3/errorhandler"
	"golang-gin3/schema"
	"golang-gin3/src/repository"
)

type TweetService interface {
	Create(payload *dto.TweetCreateDto) (*dto.TweetResponse, error)
	FindAll() ([]dto.TweetResponse, error)
	FindById(id string) (*dto.TweetResponse, error)
	Update(payload *dto.UpdateTweet, id string) (*dto.TweetResponse, error)
	Delete(id string) error
}

type tweetService struct {
	tweetRepo repository.TweetInterface
}

func NewTweetService(tweetRepo repository.TweetInterface) *tweetService {
	return &tweetService{tweetRepo}
}

func (s *tweetService) Create(payload *dto.TweetCreateDto) (*dto.TweetResponse, error) {
	tweet := schema.Tweet{
		Title:   payload.Title,
		Caption: payload.Caption,
		Image:   payload.Image,
		UserId:  payload.UserId,
	}

	err := s.tweetRepo.Create(&tweet)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	newTweet := dto.TweetResponse{
		Id:        tweet.Id,
		Title:     tweet.Title,
		Caption:   tweet.Caption,
		Image:     tweet.Caption,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}

	return &newTweet, nil
}

func (s *tweetService) FindAll() ([]dto.TweetResponse, error) {
	tweets, err := s.tweetRepo.FindAll()

	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	var tweetResponse []dto.TweetResponse

	for _, tweet := range tweets {
		response := dto.TweetResponse{
			Id:      tweet.Id,
			Title:   tweet.Title,
			Caption: tweet.Caption,
			Image:   tweet.Image,
			User: dto.UserTweet{
				Name:  tweet.User.Id,
				Email: tweet.User.Email,
			},
			CreatedAt: tweet.CreatedAt,
			UpdatedAt: tweet.UpdatedAt,
		}

		tweetResponse = append(tweetResponse, response)
	}

	return tweetResponse, nil
}

func (s *tweetService) FindById(id string) (*dto.TweetResponse, error) {
	currentTweet, err := s.tweetRepo.FindById(id)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	tweet := dto.TweetResponse{
		Id:      currentTweet.Id,
		Title:   currentTweet.Title,
		Caption: currentTweet.Caption,
		Image:   currentTweet.Image,
		User: dto.UserTweet{
			Name:  currentTweet.User.Name,
			Email: currentTweet.User.Email,
		},
		CreatedAt: currentTweet.CreatedAt,
		UpdatedAt: currentTweet.UpdatedAt,
	}

	return &tweet, nil
}

func (s *tweetService) Update(payload *dto.UpdateTweet, id string) (*dto.TweetResponse, error) {
	tweet, err := s.tweetRepo.FindById(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Not Found Error"}
	}

	tweet.Title = payload.Title
	tweet.Caption = payload.Caption
	tweet.Image = payload.Image

	err = s.tweetRepo.Update(tweet)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := dto.TweetResponse{
		Id:      tweet.Id,
		Title:   tweet.Title,
		Caption: tweet.Caption,
		Image:   tweet.Image,
	}

	return &response, nil
}

func (s *tweetService) Delete(id string) error {
	_, err := s.tweetRepo.FindById(id)

	if err != nil {
		return &errorhandler.NotFoundError{Message: "Not found"}
	}

	err = s.tweetRepo.Delete(id)

	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
