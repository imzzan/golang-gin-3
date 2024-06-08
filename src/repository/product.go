package repository

import (
	"golang-gin3/schema"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(payload *schema.Product) error
	FindAll() (*[]schema.Product, error)
	FindById(id string) (schema.Product, error)
	FindByUserId(userId string) ([]*schema.Product, error)
	Update(payload *schema.Product) error
	Delete(id string) error
}

type productRepository struct {
	Db *gorm.DB
}

func NewProductRepository(Db *gorm.DB) *productRepository {
	return &productRepository{Db}
}

func (s *productRepository) Create(payload *schema.Product) error {
	return s.Db.Create(&payload).Error
}

func (s *productRepository) FindAll() (*[]schema.Product, error) {
	var product []schema.Product
	err := s.Db.Preload("User").Find(&product).Error
	return &product, err
}

func (s *productRepository) FindById(id string) (schema.Product, error) {
	var product schema.Product
	err := s.Db.Preload("User").First(&product, "id = ?", id).Error
	return product, err
}

func (s *productRepository) FindByUserId(userId string) ([]*schema.Product, error) {
	var product []*schema.Product
	err := s.Db.Preload("User").Find(&product, "user_id = ?", userId).Error
	return product, err
}

func (s *productRepository) Update(payload *schema.Product) error {
	return s.Db.Omit("User").Save(&payload).Error
}

func (s *productRepository) Delete(id string) error {
	var product schema.Product
	return s.Db.Delete(&product, "id = ?", id).Error
}
