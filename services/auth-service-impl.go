package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/helper"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
)

type AuthServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewServiceAuth(userRepository repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *AuthServiceImpl) DoLogin(request request.LoginRequest) (*response.LoginResponse, error) {
	user, err := service.UserRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	return helper.ToLoginResponse(user), nil
}
func (service *AuthServiceImpl) DoRegister(request *request.UserRequest) (*response.UserResponse, error) {
	user, err := service.UserRepository.Create(&domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Username:  request.Username,
		Password:  request.Password,
	})
	if err != nil {
		return nil, err
	}
	return helper.ToUserResponse(user), nil
}
