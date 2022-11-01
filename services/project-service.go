package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
)

type ProjectService interface {
	FindByIdProject(projectcode string) (web.ProjectResponse, error)
	FindAllProjects() ([]web.ProjectResponse, error)
}
