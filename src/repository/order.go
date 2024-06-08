package repository

import (
	"golang-gin3/schema"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(payload schema.Order) error
	FindAll() ([]schema.Order, error)
	FindById(id string) (*schema.Order, error)
	Update(payload *schema.Order) error
	Delete(id string) error
}

type orderRepository struct {
	Db *gorm.DB
}

func NewRepositoryOrder(db *gorm.DB) *orderRepository {
	return &orderRepository{Db: db}
}

func (r *orderRepository) Create(payload schema.Order) error {
	return r.Db.Create(&payload).Error
}

func (r *orderRepository) FindAll() ([]schema.Order, error) {
	var orders []schema.Order
	err := r.Db.Preload("User").Preload("Product").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) FindById(id string) (*schema.Order, error) {
	var order schema.Order
	err := r.Db.Preload("User").Preload("Product").Find(&order, "id=?", id).Error
	return &order, err
}

func (r *orderRepository) Update(payload *schema.Order) error {
	err := r.Db.Omit("User").Save(&payload).Error
	return err
}

func (r *orderRepository) Delete(id string) error {
	var order schema.Order
	err := r.Db.Delete(&order, "id=?", id).Error
	return err
}
