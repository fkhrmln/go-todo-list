package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          string    `gorm:"column:id;type:uuid;primaryKey"`
	UserID      string    `gorm:"column:user_id"`
	Title       string    `gorm:"column:title;type:varchar(100);not null"`
	Body        string    `gorm:"column:body;type:text"`
	Deadline    int       `gorm:"column:deadline;type:integer;not null"`
	IsCompleted bool      `gorm:"column:is_completed;type:boolean;default:false;not null"`
	User        User      `gorm:"foreignKey:user_id;references:id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (todo *Todo) BeforeCreate(db *gorm.DB) error {
	todo.ID = uuid.New().String()

	return nil
}
