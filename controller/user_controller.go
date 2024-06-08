package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	FindUserById(c *fiber.Ctx) error
}
