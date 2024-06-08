package database

import "go-todo-list/entity"

func Migration() {
	db := GetConnection()

	db.Migrator().AutoMigrate(&entity.User{})
}
