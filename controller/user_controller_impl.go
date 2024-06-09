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
	signUpRequest := request.SignUpRequest{}

	err := c.BodyParser(&signUpRequest)

	if err != nil {
		return err
	}

	userResponse, err := controller.UserService.SignUp(signUpRequest)

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
	signInRequest := request.SignInRequest{}

	err := c.BodyParser(&signInRequest)

	if err != nil {
		return err
	}

	signInResponse, err := controller.UserService.SignIn(signInRequest)

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

func (controller *UserControllerImpl) GetUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	userResponse, err := controller.UserService.GetUserById(userId)

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
