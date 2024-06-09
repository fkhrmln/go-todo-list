package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
}
