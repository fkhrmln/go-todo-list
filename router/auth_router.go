package router

import (
	"go-todo-list/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(apiRouter fiber.Router, controller controller.UserController) {
	authRouter := apiRouter.Group("/auth")

	authRouter.Post("/signup", controller.SignUp)
	authRouter.Post("/signin", controller.SignIn)
}
