package repository

import (
	"go-todo-list/entity"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (repository *TodoRepositoryImpl) Create(db *gorm.DB, todo entity.Todo) (entity.Todo, error) {
	err := db.Create(&todo).Error

	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}

func (repository *TodoRepositoryImpl) FindAll(db *gorm.DB, userId string) ([]entity.Todo, error) {
	todos := []entity.Todo{}

	err := db.Where("user_id = ?", userId).Find(&todos).Error

	if err != nil {
		return []entity.Todo{}, err
	}

	return todos, nil
}

func (repository *TodoRepositoryImpl) FindById(db *gorm.DB, userId string, todoId string) (entity.Todo, error) {
	todo := entity.Todo{}

	err := db.Where("user_id = ? AND id = ?", userId, todoId).First(&todo).Error

	if err != nil {
		return entity.Todo{}, err
	}

	return todo, err
}

func (repository *TodoRepositoryImpl) Update(db *gorm.DB, todo entity.Todo, updateTodo map[string]interface{}) (entity.Todo, error) {
	err := db.Model(&todo).Updates(updateTodo).Error

	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}

func (repository *TodoRepositoryImpl) Delete(db *gorm.DB, userId string, todoId string) error {
	todo := entity.Todo{}

	err := db.Where("user_id = ? AND id = ?", userId, todoId).Delete(&todo).Error

	if err != nil {
		return err
	}

	return nil
}
