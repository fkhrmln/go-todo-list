package controller

import "github.com/gofiber/fiber/v2"

type TodoController interface {
	CreateTodo(c *fiber.Ctx) error
	GetTodos(c *fiber.Ctx) error
	GetTodoById(c *fiber.Ctx) error
	UpdateTodo(c *fiber.Ctx) error
	DeleteTodo(c *fiber.Ctx) error
}
