package repository

import (
	"go-todo-list/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(db *gorm.DB, user entity.User) (entity.User, error) {
	err := db.Create(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindById(db *gorm.DB, userId string) (entity.User, error) {
	user := entity.User{}

	err := db.Where("id = ?", userId).Take(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(db *gorm.DB, email string) (entity.User, error) {
	user := entity.User{}

	err := db.Where("email = ?", email).Take(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
