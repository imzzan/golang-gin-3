package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tweet struct {
	Id        string `gorm:"primarykey"`
	Title     string `gorm:"type:varchar(255);not null"`
	Caption   string `gorm:"type:varchar(255);not null"`
	Image     string `gorm:"type:varchar(255);not null"`
	UserId    string
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (base *Tweet) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	tx.Statement.SetColumn("id", uuid)
	return nil
}
