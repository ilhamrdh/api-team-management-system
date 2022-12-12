package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
)

type ProjectService interface {
	CreateProject(request *request.ProjectRequest) (*response.ProjectResponse, error)
	UpdateProject(projectCode string, request *request.ProjectRequestUpdate) (*response.ProjectResponse, error)
	DeleteProject(projectCode string) error
	FindByIdProject(projctCode string) (*response.ProjectResponse, error)
	FindAllProjects() ([]*response.ProjectResponse, error)
}
