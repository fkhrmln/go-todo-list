package router

import (
	"go-todo-list/controller"
	"go-todo-list/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(apiRouter fiber.Router, controller controller.UserController) {
	authRouter := apiRouter.Group("/users")

	authRouter.Get("/:userId", middleware.VerifyToken(), controller.FindUserById)
}
