package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey"`
	Email     string    `gorm:"column:email;type:varchar(100);unique;not null"`
	Password  string    `gorm:"column:password;type:varchar(100);not null"`
	Username  string    `gorm:"column:username;type:varchar(50);not null"`
	Todos     []Todo    `gorm:"foreignKey:user_id;references:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.ID = uuid.New().String()

	return nil
}
