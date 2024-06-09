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

func (service *UserServiceImpl) SignUp(signUpRequest request.SignUpRequest) (response.UserResponse, error) {
	err := service.Validator.Struct(signUpRequest)

	if err != nil {
		return response.UserResponse{}, &exception.ValidationError{Message: err.Error()}
	}

	_, err = service.UserRepository.FindByEmail(service.DB, signUpRequest.Email)

	if err == nil {
		return response.UserResponse{}, &exception.BadRequestError{Message: "Email Already Taken"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), 14)

	if err != nil {
		return response.UserResponse{}, err
	}

	user := entity.User{
		Email:    signUpRequest.Email,
		Password: string(hashedPassword),
		Username: signUpRequest.Username,
	}

	user, err = service.UserRepository.Create(service.DB, user)

	if err != nil {
		return response.UserResponse{}, err
	}

	userResponse := helper.UserToUserResponse(user)

	return userResponse, nil
}

func (service *UserServiceImpl) SignIn(signInRequest request.SignInRequest) (response.SignInResponse, error) {
	err := service.Validator.Struct(signInRequest)

	if err != nil {
		return response.SignInResponse{}, &exception.ValidationError{Message: err.Error()}
	}

	user, err := service.UserRepository.FindByEmail(service.DB, signInRequest.Email)

	if err != nil {
		return response.SignInResponse{}, &exception.BadRequestError{Message: "User Not Registered"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequest.Password))

	if err != nil {
		return response.SignInResponse{}, &exception.BadRequestError{Message: "Wrong Password"}
	}

	token := helper.GenerateToken(user)

	signInResponse := response.SignInResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Token:    token,
	}

	return signInResponse, nil
}

func (service *UserServiceImpl) GetUserById(userId string) (response.UserResponse, error) {
	user, err := service.UserRepository.FindById(service.DB, userId)

	if err != nil {
		return response.UserResponse{}, &exception.NotFoundError{Message: "User Not Found"}
	}

	userResponse := helper.UserToUserResponse(user)

	return userResponse, nil
}
