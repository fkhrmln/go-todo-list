package router

import (
	"go-todo-list/controller"
	"go-todo-list/middleware"

	"github.com/gofiber/fiber/v2"
)

func TodoRouter(apiRouter fiber.Router, controller controller.TodoController) {
	authRouter := apiRouter.Group("/todos")

	authRouter.Post("/", middleware.VerifyToken(), controller.CreateTodo)
	authRouter.Get("/", middleware.VerifyToken(), controller.GetTodos)
	authRouter.Get("/:todoId", middleware.VerifyToken(), controller.GetTodoById)
	authRouter.Put("/:todoId", middleware.VerifyToken(), controller.UpdateTodo)
	authRouter.Delete("/:todoId", middleware.VerifyToken(), controller.DeleteTodo)
}
