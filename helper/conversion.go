package helper

import (
	"go-todo-list/dto/response"
	"go-todo-list/entity"
)

func UserToUserResponse(user entity.User) response.UserResponse {
	userResponse := response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return userResponse
}
