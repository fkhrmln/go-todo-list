package app

import (
	"go-todo-list/exception"

	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.CustomErrorHandler,
	})

	return app
}
