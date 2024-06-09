package service

import (
	"go-todo-list/dto/request"
	"go-todo-list/dto/response"
)

type TodoService interface {
	CreateTodo(createTodoRequest request.CreateTodoRequest) (response.TodoResponse, error)
	GetTodos(userId string) ([]response.TodoResponse, error)
	GetTodoById(userId string, todoId string) (response.TodoResponse, error)
	UpdateTodo(updateTodoRequest request.UpdateTodoRequest) (response.TodoResponse, error)
	DeleteTodo(userId string, todoId string) error
}
