package repository

import (
	"go-todo-list/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user entity.User) entity.User
	FindById(db *gorm.DB, userId string) (entity.User, error)
	FindByUsername(db *gorm.DB, username string) (entity.User, error)
}
