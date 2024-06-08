package repository

import (
	"golang-gin3/schema"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(payload *schema.User) error
	FindEmail(email string) (*schema.User, error)
	FindById(id string) (*schema.User, error)
}

type userReposirtory struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userReposirtory {
	return &userReposirtory{db}
}

func (r *userReposirtory) Create(payload *schema.User) error {
	return r.Db.Create(&payload).Error
}

func (r *userReposirtory) FindEmail(email string) (*schema.User, error) {
	var user schema.User
	err := r.Db.Find(&user, "email=?", email).Error
	return &user, err
}

func (r *userReposirtory) FindById(id string) (*schema.User, error) {
	var user schema.User
	err := r.Db.Find(&user, "id=?", id).Error
	return &user, err
}
