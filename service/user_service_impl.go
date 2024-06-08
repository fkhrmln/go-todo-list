package service

import (
	"go-todo-list/dto/request"
	"go-todo-list/dto/response"
	"go-todo-list/entity"
	"go-todo-list/exception"
	"go-todo-list/helper"
	"go-todo-list/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validator      *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *gorm.DB, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validator:      validator,
	}
}

func (service *UserServiceImpl) SignUp(authRequest request.AuthRequest) (response.UserResponse, error) {
	err := service.Validator.Struct(authRequest)

	if err != nil {
		return response.UserResponse{}, &exception.ValidationError{Message: err.Error()}
	}

	user, _ := service.UserRepository.FindByUsername(service.DB, authRequest.Username)

	if user.ID != "" {
		return response.UserResponse{}, &exception.BadRequestError{Message: "Username Already Taken"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), 14)

	if err != nil {
		return response.UserResponse{}, err
	}

	user = entity.User{
		Username: authRequest.Username,
		Password: string(hashedPassword),
	}

	user = service.UserRepository.Create(service.DB, user)

	userResponse := helper.UserToUserResponse(user)

	return userResponse, nil
}

func (service *UserServiceImpl) SignIn(authRequest request.AuthRequest) (response.SignInResponse, error) {
	err := service.Validator.Struct(authRequest)

	if err != nil {
		return response.SignInResponse{}, &exception.ValidationError{Message: err.Error()}
	}

	user, _ := service.UserRepository.FindByUsername(service.DB, authRequest.Username)

	if user.ID == "" {
		return response.SignInResponse{}, &exception.BadRequestError{Message: "User Not Registered"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password))

	if err != nil {
		return response.SignInResponse{}, &exception.BadRequestError{Message: "Wrong Password"}
	}

	token := helper.GenerateToken(user)

	signInResponse := response.SignInResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}

	return signInResponse, nil
}

func (service *UserServiceImpl) FindUserById(userId string) (response.UserResponse, error) {
	user, err := service.UserRepository.FindById(service.DB, userId)

	if err != nil {
		return response.UserResponse{}, &exception.NotFoundError{Message: "User Not Found"}
	}

	userResponse := helper.UserToUserResponse(user)

	return userResponse, nil
}
