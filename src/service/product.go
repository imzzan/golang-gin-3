package service

import (
	"golang-gin3/dto"
	"golang-gin3/errorhandler"
	"golang-gin3/schema"
	"golang-gin3/src/repository"
)

type ProductService interface {
	Create(payload *dto.ProductDto) (*dto.ProductResponse, error)
	FindAll() ([]dto.ProductResponse, error)
	FindById(id string) (*dto.ProductResponse, error)
	FindByUserId(userId string) ([]dto.ProductResponse, error)
	Update(payload *dto.UpdateProductDto, id string) (*dto.UpdateProductResponse, error)
	Delete(id string) error
}

type productService struct {
	productRepo    repository.ProductRepository
	userRepository repository.UserRepository
}

func NewProductService(productRepo repository.ProductRepository, userRepository repository.UserRepository) *productService {
	return &productService{productRepo, userRepository}
}

func (s *productService) Create(payload *dto.ProductDto) (*dto.ProductResponse, error) {
	product := schema.Product{
		Name:   payload.Name,
		Price:  payload.Price,
		Image:  payload.Image.Filename,
		UserId: payload.UserId,
	}

	user, err := s.userRepository.FindById(payload.UserId)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	err = s.productRepo.Create(&product)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := dto.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Image: product.Image,
		User: dto.UserProduct{
			Name:  user.Name,
			Email: user.Email,
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return &response, nil
}

func (s *productService) FindAll() ([]dto.ProductResponse, error) {
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	var productsResponse []dto.ProductResponse

	for _, product := range *products {
		response := dto.ProductResponse{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
			Image: product.Image,
			User: dto.UserProduct{
				Name:  product.User.Name,
				Email: product.User.Email,
			},
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}

		productsResponse = append(productsResponse, response)
	}

	return productsResponse, nil
}

func (s *productService) FindById(id string) (*dto.ProductResponse, error) {
	product, err := s.productRepo.FindById(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Product not found"}
	}

	response := dto.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Image: product.Image,
		User: dto.UserProduct{
			Name:  product.User.Name,
			Email: product.User.Email,
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return &response, nil
}

func (s *productService) FindByUserId(userId string) ([]dto.ProductResponse, error) {
	products, err := s.productRepo.FindByUserId(userId)

	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: err.Error()}
	}

	var productResponse []dto.ProductResponse

	for _, item := range products {
		response := dto.ProductResponse{
			Id:    item.Id,
			Name:  item.Name,
			Price: item.Price,
			Image: item.Image,
			User: dto.UserProduct{
				Name:  item.User.Name,
				Email: item.User.Email,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		productResponse = append(productResponse, response)
	}

	return productResponse, nil
}

func (s *productService) Update(payload *dto.UpdateProductDto, id string) (*dto.UpdateProductResponse, error) {
	product, err := s.productRepo.FindById(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Product not found"}
	}

	product.Name = payload.Name
	product.Price = payload.Price

	err = s.productRepo.Update(&product)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := dto.UpdateProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Image: product.Image,
	}

	return &response, err
}

func (s *productService) Delete(id string) error {
	_, err := s.productRepo.FindById(id)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "Product not found"}
	}

	err = s.productRepo.Delete(id)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
