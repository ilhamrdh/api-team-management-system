package repositories

import "github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"

type ProjectRepository interface {
	FindByIdProject(projectcode string) (domain.Project, error)
	FindAllProjects() ([]domain.Project, error)
}
