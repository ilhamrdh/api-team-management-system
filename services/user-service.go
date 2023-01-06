package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
)

type UserService interface {
	FindById(id int) (*response.UserResponse, error)
	FindByUsername(username string) (*response.UserResponse, error)
	Update(id int, request *request.UserRequest) (*response.UserResponse, error)
	Delete(id int) error
}
