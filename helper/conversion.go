package helper

import (
	"go-todo-list/dto/response"
	"go-todo-list/entity"
)

func UserToUserResponse(user entity.User) response.UserResponse {
	userResponse := response.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}

	return userResponse
}

func TodoToTodoResponse(todo entity.Todo) response.TodoResponse {
	todoResponse := response.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Body:        todo.Body,
		Deadline:    todo.Deadline,
		IsCompleted: todo.IsCompleted,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}

	return todoResponse
}
