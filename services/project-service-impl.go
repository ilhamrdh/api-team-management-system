package services

import (
	"time"

	"github.com/IlhamRamadhan-IR/api-team-management-system/helper"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
	"github.com/go-playground/validator/v10"
)

type ProjectServiceImpl struct {
	ProjectRepository repositories.ProjectRepository
	Validate          *validator.Validate
}

func NewServiceProject(projectRepository repositories.ProjectRepository, validate *validator.Validate) ProjectService {
	return &ProjectServiceImpl{
		ProjectRepository: projectRepository,
		Validate:          validate,
	}
}

func (service *ProjectServiceImpl) FindByIdProject(projctCode string) (*response.ProjectResponse, error) {
	project, err := service.ProjectRepository.FindByIdProject(projctCode)
	if err != nil {
		return nil, err
	}
	return helper.ToProjectResponse(project), nil
}
func (service *ProjectServiceImpl) FindAllProjects() ([]*response.ProjectResponse, error) {
	projects, err := service.ProjectRepository.FindAllProjects()
	if err != nil {
		return nil, err
	}
	return helper.ToProjectResponses(projects), nil
}

func (service *ProjectServiceImpl) CreateProject(request *request.ProjectRequest) (*response.ProjectResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return nil, err
	}
	
	project, err := service.ProjectRepository.CreateProject(&domain.Project{
		ProjectCode: request.ProjectCode,
		ProjectName: request.ProjectName,
		Deadline:    request.Deadline,
		Status:      request.Status,
	})
	if err != nil {
		return nil, err
	}

	if _ , err := time.Parse("2006-01-02",request.Deadline); err != nil {
		return nil, err
	}

	return helper.ToProjectResponse(project), nil
}
func (service *ProjectServiceImpl) UpdateProject(projectCode string, request *request.ProjectRequestUpdate) (*response.ProjectResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return nil, err
	}
	_, err := service.ProjectRepository.FindByIdProject(projectCode)
	if err != nil {
		return nil, err
	}
	project, err := service.ProjectRepository.UpdateProject(projectCode, &domain.Project{
		ProjectName: request.ProjectName,
		Deadline:    request.Deadline,
		Status:      request.Status,
	})
	if err != nil {
		return nil, err
	}
	if _ , err := time.Parse("2006-01-02",request.Deadline); err != nil {
		return nil, err
	}
	return helper.ToProjectResponse(project), nil
}
func (service *ProjectServiceImpl) DeleteProject(projectCode string) error {
	err := service.ProjectRepository.DeleteProject(projectCode)
	if err != nil {
		return err
	}
	return nil
}
