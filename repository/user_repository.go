package repository

import (
	"go-todo-list/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user entity.User) (entity.User, error)
	FindById(db *gorm.DB, userId string) (entity.User, error)
	FindByEmail(db *gorm.DB, email string) (entity.User, error)
}
