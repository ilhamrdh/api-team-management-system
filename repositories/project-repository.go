package repositories

import "github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"

type ProjectRepository interface {
	CreateProject(project *domain.Project) (*domain.Project, error)
	UpdateProject(projectCode string, project *domain.Project) (*domain.Project, error)
	DeleteProject(projectCode string) error
	FindByIdProject(projectCode string) (*domain.Project, error)
	FindAllProjects() ([]*domain.Project, error)
}
