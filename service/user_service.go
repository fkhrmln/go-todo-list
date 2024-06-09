package service

import (
	"go-todo-list/dto/request"
	"go-todo-list/dto/response"
)

type UserService interface {
	SignUp(user request.SignUpRequest) (response.UserResponse, error)
	SignIn(user request.SignInRequest) (response.SignInResponse, error)
	GetUserById(userId string) (response.UserResponse, error)
}
