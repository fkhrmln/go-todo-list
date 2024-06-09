package main

import (
	"fmt"
	"go-todo-list/app"
	"go-todo-list/controller"
	"go-todo-list/database"
	"go-todo-list/repository"
	"go-todo-list/router"
	"go-todo-list/service"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	db := database.GetConnection()

	database.Migration()

	validator := validator.New()

	app := app.NewApp()

	userRepository := repository.NewUserRepository()

	userService := service.NewUserService(userRepository, db, validator)

	userController := controller.NewUserController(userService)

	todoRepository := repository.NewTodoRepository()

	todoService := service.NewTodoService(todoRepository, db, validator)

	todoController := controller.NewTodoController(todoService)

	apiRouter := app.Group("/api/v1")

	router.AuthRouter(apiRouter, userController)

	router.UserRouter(apiRouter, userController)

	router.TodoRouter(apiRouter, todoController)

	err = app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))

	if err != nil {
		panic(err)
	}
}
