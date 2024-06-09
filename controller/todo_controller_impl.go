package controller

import (
	"go-todo-list/dto/request"
	"go-todo-list/dto/response"
	"go-todo-list/helper"
	"go-todo-list/service"

	"github.com/gofiber/fiber/v2"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (controller *TodoControllerImpl) CreateTodo(c *fiber.Ctx) error {
	userId := helper.GetUserIdFromToken(c)

	createTodoRequest := request.CreateTodoRequest{}

	err := c.BodyParser(&createTodoRequest)

	if err != nil {
		return err
	}

	createTodoRequest.UserID = userId

	todoResponse, err := controller.TodoService.CreateTodo(createTodoRequest)

	if err != nil {
		return err
	}

	response := response.Response{
		Status:  "Created",
		Code:    fiber.StatusCreated,
		Message: "Todo Created Successfully",
		Data:    todoResponse,
	}

	return c.Status(201).JSON(response)
}

func (controller *TodoControllerImpl) GetTodos(c *fiber.Ctx) error {
	userId := helper.GetUserIdFromToken(c)

	todosResponse, err := controller.TodoService.GetTodos(userId)

	if err != nil {
		return err
	}

	response := response.Response{
		Status:  "OK",
		Code:    fiber.StatusOK,
		Message: "",
		Data:    todosResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *TodoControllerImpl) GetTodoById(c *fiber.Ctx) error {
	userId := helper.GetUserIdFromToken(c)

	todoId := c.Params("todoId")

	todoResponse, err := controller.TodoService.GetTodoById(userId, todoId)

	if err != nil {
		return err
	}

	response := response.Response{
		Status:  "OK",
		Code:    fiber.StatusOK,
		Message: "",
		Data:    todoResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *TodoControllerImpl) UpdateTodo(c *fiber.Ctx) error {
	userId := helper.GetUserIdFromToken(c)

	todoId := c.Params("todoId")

	updateTodoRequest := request.UpdateTodoRequest{}

	err := c.BodyParser(&updateTodoRequest)

	if err != nil {
		return err
	}

	updateTodoRequest.ID = todoId
	updateTodoRequest.UserID = userId

	todoResponse, err := controller.TodoService.UpdateTodo(updateTodoRequest)

	if err != nil {
		return err
	}

	response := response.Response{
		Status:  "OK",
		Code:    fiber.StatusOK,
		Message: "Todo Updated Successfully",
		Data:    todoResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *TodoControllerImpl) DeleteTodo(c *fiber.Ctx) error {
	userId := helper.GetUserIdFromToken(c)

	todoId := c.Params("todoId")

	err := controller.TodoService.DeleteTodo(userId, todoId)

	if err != nil {
		return err
	}

	response := response.Response{
		Status:  "OK",
		Code:    fiber.StatusOK,
		Message: "Todo Deleted Successfully",
		Data:    nil,
	}

	return c.Status(200).JSON(response)
}
