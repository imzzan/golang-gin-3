package service

import (
	"golang-gin3/dto"
	"golang-gin3/errorhandler"
	"golang-gin3/helper"
	"golang-gin3/schema"
	"golang-gin3/src/repository"
)

type UserService interface {
	Create(payload *dto.UserDto) (*dto.UserResponse, error)
	Login(payload *dto.LoginDto) (*dto.LoginResponse, error)
	GetMe(id string) (*dto.UserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Create(payload *dto.UserDto) (*dto.UserResponse, error) {

	currentUser, _ := s.userRepository.FindEmail(payload.Email)
	if payload.Email == currentUser.Email {
		return nil, &errorhandler.ConflictError{Message: "Email Already Exists"}
	}

	hashPassword, err := helper.HashPassword(payload.Password)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := schema.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashPassword,
	}

	err = s.userRepository.Create(&user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	newuser := dto.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
	}

	return &newuser, nil
}

func (s *userService) Login(payload *dto.LoginDto) (*dto.LoginResponse, error) {
	currentUser, _ := s.userRepository.FindEmail(payload.Email)
	if currentUser.Email != payload.Email {
		return nil, &errorhandler.BadRequestError{Message: "Email not found"}
	}

	err := helper.MatchPasword(payload.Password, currentUser.Password)
	if err != nil {
		return nil, &errorhandler.BadRequestError{Message: "Password is incorrect"}
	}

	token, err := helper.GenerateToken(currentUser)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := dto.LoginResponse{
		Token: token,
	}

	return &response, nil
}

func (s *userService) GetMe(id string) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "User not found"}
	}

	response := dto.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return &response, err
}
