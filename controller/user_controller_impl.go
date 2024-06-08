package controller

import (
	"go-todo-list/dto/request"
	"go-todo-list/dto/response"
	"go-todo-list/service"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) SignUp(c *fiber.Ctx) error {
	AuthRequest := request.AuthRequest{}

	err := c.BodyParser(&AuthRequest)

	if err != nil {
		return err
	}

	userResponse, err := controller.UserService.SignUp(AuthRequest)

	if err != nil {
		return err
	}

	response := response.Response{
		Status:  "Created",
		Code:    fiber.StatusCreated,
		Message: "SignUp Successfully",
		Data:    userResponse,
	}

	return c.Status(201).JSON(response)
}

func (controller *UserControllerImpl) SignIn(c *fiber.Ctx) error {
	AuthRequest := request.AuthRequest{}

	err := c.BodyParser(&AuthRequest)

	if err != nil {
		return err
	}

	signInResponse, err := controller.UserService.SignIn(AuthRequest)

	if err != nil {
		return err
	}

	response := response.Response{
		Status:  "OK",
		Code:    fiber.StatusOK,
		Message: "SignIn Successfully",
		Data:    signInResponse,
	}

	c.Cookie(&fiber.Cookie{
		Name:  "JWT",
		Value: signInResponse.Token,
	})

	return c.Status(200).JSON(response)
}

func (controller *UserControllerImpl) FindUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	userResponse, err := controller.UserService.FindUserById(userId)

	if err != nil {
		return err
	}

	response := response.Response{
		Status:  "OK",
		Code:    fiber.StatusOK,
		Message: "",
		Data:    userResponse,
	}

	return c.Status(200).JSON(response)
}
