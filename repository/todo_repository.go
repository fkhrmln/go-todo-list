package repository

import (
	"go-todo-list/entity"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(db *gorm.DB, todo entity.Todo) (entity.Todo, error)
	FindAll(db *gorm.DB, userId string) ([]entity.Todo, error)
	FindById(db *gorm.DB, userId string, todoId string) (entity.Todo, error)
	Update(db *gorm.DB, todo entity.Todo, updateTodo map[string]interface{}) (entity.Todo, error)
	Delete(db *gorm.DB, userId string, todoId string) error
}
