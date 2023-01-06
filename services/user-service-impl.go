package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/helper"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewServiceUser(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) FindById(id int) (*response.UserResponse, error) {
	user, err := service.UserRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return helper.ToUserResponse(user), nil
}
func (service *UserServiceImpl) FindByUsername(username string) (*response.UserResponse, error) {
	user, err := service.UserRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return helper.ToUserResponse(user), nil
}
func (service *UserServiceImpl) Update(id int, request *request.UserRequest) (*response.UserResponse, error) {
	user, err := service.UserRepository.Update(id, &domain.User{
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
func (service *UserServiceImpl) Delete(id int) error {
	err := service.UserRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
