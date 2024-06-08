package service

import (
	"golang-gin3/dto"
	"golang-gin3/errorhandler"
	"golang-gin3/schema"
	"golang-gin3/src/repository"
)

type OrderService interface {
	Create(paylaod dto.OrderDto) error
	FindAll() ([]dto.OrderResponse, error)
	FindById(id string) (*dto.OrderResponse, error)
	Update(id string, payload dto.OrderDto) error
	Delete(id string) error
}

type orderService struct {
	orseeRepository repository.OrderRepository
}

func NewOrderService(orseeRepository repository.OrderRepository) *orderService {
	return &orderService{orseeRepository}
}

func (s *orderService) Create(paylaod dto.OrderDto) error {
	order := schema.Order{
		UserId:    paylaod.UserId,
		ProductId: paylaod.ProductId,
	}
	err := s.orseeRepository.Create(order)
	if err != nil {
		return &errorhandler.BadRequestError{Message: "Gagal Input data"}
	}

	return nil
}
func (s *orderService) FindAll() ([]dto.OrderResponse, error) {
	row, err := s.orseeRepository.FindAll()
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: "Internal Server Error"}
	}

	var orders []dto.OrderResponse

	for _, r := range row {
		response := dto.OrderResponse{
			Id: r.Id,
			User: dto.UserResponse{
				Id:   r.User.Id,
				Name: r.User.Name,
			},
			Product: dto.ProductResponse{
				Id:   r.Product.Id,
				Name: r.Product.Name,
			},
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		}

		orders = append(orders, response)
	}

	return orders, nil
}

func (s *orderService) FindById(id string) (*dto.OrderResponse, error) {
	order, err := s.orseeRepository.FindById(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	response := dto.OrderResponse{
		Id:     order.Id,
		IsPaid: order.IsPaid,
		User: dto.UserResponse{
			Id:    order.User.Id,
			Name:  order.User.Name,
			Email: order.User.Email,
		},
		Product: dto.ProductResponse{
			Id:    order.Product.Id,
			Name:  order.Product.Name,
			Price: order.Product.Price,
		},
	}

	return &response, err
}

func (s *orderService) Update(id string, payload dto.OrderDto) error {

	order, err := s.orseeRepository.FindById(id)
	if err != nil {
		return &errorhandler.BadRequestError{Message: "Data Not Found"}
	}

	order.IsPaid = true
	order.ProductId = payload.ProductId
	order.UserId = payload.UserId
	err = s.orseeRepository.Update(order)
	if err != nil {
		return &errorhandler.InternalServerError{Message: "Internal Server Error"}
	}

	return nil
}

func (s *orderService) Delete(id string) error {
	_, err := s.orseeRepository.FindById(id)
	if err != nil {
		return &errorhandler.BadRequestError{Message: "Data Not Found"}
	}

	err = s.orseeRepository.Delete(id)
	if err != nil {
		return &errorhandler.BadRequestError{Message: "Internal Server"}
	}

	return nil
}
