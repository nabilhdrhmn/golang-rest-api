package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `gorm:"unique;" json:"phone_number"`
	Address     string    `json:"address"`
	PIN         string    `json:"pin"`
	CreatedAt   time.Time `json:"created_at"`
	Balance     int64     `json:"balance"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	u.ID = uuid.New()
	return nil
}
