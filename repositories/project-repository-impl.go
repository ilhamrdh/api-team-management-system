package repositories

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
	DB *gorm.DB
}

func NewRepositoryProject(db *gorm.DB) ProjectRepository {
	return &ProjectRepositoryImpl{
		DB: db,
	}
}
func (repository *ProjectRepositoryImpl) FindByIdProject(projectcode string) (domain.Project, error) {
	var project domain.Project
	err := repository.DB.Where("project_code = ?", projectcode).First(&project).Error
	return project, err
}
func (repository *ProjectRepositoryImpl) FindAllProjects() ([]domain.Project, error) {
	var projects []domain.Project
	err := repository.DB.Find(&projects).Error
	return projects, err
}
