package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
)

type AuthService interface {
	DoRegister(request *request.UserRequest) (*response.UserResponse, error)
	DoLogin(request request.LoginRequest) (*response.LoginResponse, error)
}
