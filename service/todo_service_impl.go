package service

import (
	"go-todo-list/dto/request"
	"go-todo-list/dto/response"
	"go-todo-list/entity"
	"go-todo-list/exception"
	"go-todo-list/helper"
	"go-todo-list/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	DB             *gorm.DB
	Validator      *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository, db *gorm.DB, validator *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		DB:             db,
		Validator:      validator,
	}
}

func (service *TodoServiceImpl) CreateTodo(createTodoRequest request.CreateTodoRequest) (response.TodoResponse, error) {
	err := service.Validator.Struct(createTodoRequest)

	if err != nil {
		return response.TodoResponse{}, &exception.ValidationError{Message: err.Error()}
	}

	todo := entity.Todo{
		UserID:   createTodoRequest.UserID,
		Title:    createTodoRequest.Title,
		Body:     createTodoRequest.Body,
		Deadline: createTodoRequest.Deadline,
	}

	todo, err = service.TodoRepository.Create(service.DB, todo)

	if err != nil {
		return response.TodoResponse{}, err
	}

	todoResponse := helper.TodoToTodoResponse(todo)

	return todoResponse, nil
}

func (service *TodoServiceImpl) GetTodos(userId string) ([]response.TodoResponse, error) {
	todos, err := service.TodoRepository.FindAll(service.DB, userId)

	if err != nil {
		return []response.TodoResponse{}, nil
	}

	todosResponse := []response.TodoResponse{}

	for _, todo := range todos {
		todosResponse = append(todosResponse, helper.TodoToTodoResponse(todo))
	}

	return todosResponse, nil
}

func (service *TodoServiceImpl) GetTodoById(userId string, todoId string) (response.TodoResponse, error) {
	todo, err := service.TodoRepository.FindById(service.DB, userId, todoId)

	if err != nil {
		return response.TodoResponse{}, &exception.NotFoundError{Message: "Todo Not Found"}
	}

	todoResponse := helper.TodoToTodoResponse(todo)

	return todoResponse, nil
}

func (service *TodoServiceImpl) UpdateTodo(updateTodoRequest request.UpdateTodoRequest) (response.TodoResponse, error) {
	err := service.Validator.Struct(updateTodoRequest)

	if err != nil {
		return response.TodoResponse{}, &exception.ValidationError{Message: err.Error()}
	}

	todo, err := service.TodoRepository.FindById(service.DB, updateTodoRequest.UserID, updateTodoRequest.ID)

	if err != nil {
		return response.TodoResponse{}, &exception.NotFoundError{Message: "Todo Not Found"}
	}

	updateTodo := helper.CheckValueUpdateTodoRequest(updateTodoRequest)

	todo, err = service.TodoRepository.Update(service.DB, todo, updateTodo)

	if err != nil {
		return response.TodoResponse{}, err
	}

	todoResponse := helper.TodoToTodoResponse(todo)

	return todoResponse, nil
}

func (service *TodoServiceImpl) DeleteTodo(userId string, todoId string) error {
	err := service.TodoRepository.Delete(service.DB, userId, todoId)

	if err != nil {
		return &exception.NotFoundError{Message: "Todo Not Found"}
	}

	return nil
}
