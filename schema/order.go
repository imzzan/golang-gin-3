package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Id        string `gorm:"primaryKey"`
	UserId    string
	ProductId string
	User      User
	Product   Product
	IsPaid    bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (base *Order) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	tx.Statement.SetColumn("id", uuid)
	return nil
}
