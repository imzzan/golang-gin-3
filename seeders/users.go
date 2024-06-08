package seeders

import (
	"golang-gin3/schema"
	"time"

	"gorm.io/gorm"
)

func CreateUserSeeders(Db *gorm.DB) {
	data := schema.User{
		Name:      "admin",
		Email:     "admin.gmail.com",
		Password:  "admin123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	Db.Create(&data)
}
