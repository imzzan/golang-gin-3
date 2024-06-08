package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string    `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(200);not null"`
	Email     string    `gorm:"unique;type:varchar(200);not null"`
	Password  string    `gorm:"type:varchar(200);not null"`
	Product   []Product `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tweet     []Tweet   `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Order     []Order   `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (base *User) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	tx.Statement.SetColumn("id", uuid)
	return nil
}
