package service

import (
	"go-todo-list/dto/request"
	"go-todo-list/dto/response"
)

type UserService interface {
	SignUp(user request.AuthRequest) (response.UserResponse, error)
	SignIn(user request.AuthRequest) (response.SignInResponse, error)
	FindUserById(userId string) (response.UserResponse, error)
}
