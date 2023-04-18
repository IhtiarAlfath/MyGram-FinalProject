package service

import (
	"kominfo/h8-myGram-finalproject/dto"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/pkg/helpers"
	"kominfo/h8-myGram-finalproject/repository/user_repository"
	"net/http"
)

type userService struct {
	userRepo user_repository.UserRepository
}

type UserService interface {
	CreateNewUser(newUserRequest dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(payload dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) CreateNewUser(newUserRequest dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUserRequest)

	if err != nil {
		return nil, err
	}

	err = helpers.ValidateAgeUser(newUserRequest.Age)

	if err != nil {
		return nil, err
	}
	
	payload := entity.User{
		Username: newUserRequest.Username,
		Email:    newUserRequest.Email,
		Password: newUserRequest.Password,
		Age: newUserRequest.Age,
	}

	err = payload.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.Register(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "registered successfully",
	}

	return &response, nil
}

func (u *userService) Login(payload dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr){
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetuserByEmail(payload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound{
			return nil, errs.NewNotFoundError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(payload.Password)

	if !isValidPassword{
		return nil, errs.NewNotFoundError("invalid email/password")
	}

	response := dto.LoginResponse{
		Result: "success",
		StatusCode: http.StatusOK,
		Message: "login successfully",
		Data: dto.TokenResponse{
			Token: user.GenerateToke(),
		},
	}

	return &response, nil
}