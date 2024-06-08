package config

import (
	"fmt"
	"golang-gin3/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func LoadDb() {
	dsn := fmt.Sprintf("%v", ENV.DB_URL)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&schema.User{}, &schema.Product{}, &schema.Tweet{}, &schema.Order{})

	// seeders.CreateUserSeeders(db)

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database Connection")
	Db = db
}
