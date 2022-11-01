package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/helper"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
)

type ProjectServiceImpl struct {
	ProjectRepository repositories.ProjectRepository
}

func (service *ProjectServiceImpl) FindByIdProject(projectcode string) (web.ProjectResponse, error) {
	project, err := service.ProjectRepository.FindByIdProject(projectcode)
	return helper.ToProjectResponse(project), err
}
func (service *ProjectServiceImpl) FindAllProjects() ([]web.ProjectResponse, error) {
	projects, err := service.ProjectRepository.FindAllProjects()
	return helper.ToProjectResponses(projects), err
}
