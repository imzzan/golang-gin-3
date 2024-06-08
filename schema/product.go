package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Price     int
	Image     string
	UserId    string
	User      *User
	Order     []Order `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (base *Product) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	tx.Statement.SetColumn("id", uuid)
	return nil
}
