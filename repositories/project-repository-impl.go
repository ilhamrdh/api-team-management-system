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

func (repository *ProjectRepositoryImpl) FindByIdProject(projectCode string) (*domain.Project, error) {
	var project *domain.Project
	err := repository.DB.Where("project_code = ?", projectCode).First(&project).Error
	if err != nil {
		return nil, err
	}
	return project, nil
}
func (repository *ProjectRepositoryImpl) FindAllProjects() ([]*domain.Project, error) {
	var projects []*domain.Project
	err := repository.DB.Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (repository *ProjectRepositoryImpl) CreateProject(project *domain.Project) (*domain.Project, error) {
	err := repository.DB.Exec("INSERT INTO project (project_code,project_name,deadline,status) VALUES (?,?,?,?)",
		project.ProjectCode,
		project.ProjectName,
		project.Deadline,
		project.Status,
	).Error
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (repository *ProjectRepositoryImpl) UpdateProject(projectCode string, project *domain.Project) (*domain.Project, error) {
	err := repository.DB.Exec("UPDATE project SET project_name=?,deadline=?,status=? WHERE project_code=?",
		project.ProjectName,
		project.Deadline,
		project.Status,
		projectCode,
	).Error
	if err != nil {
		return nil, err
	}
	return project, nil
}
func (repository *ProjectRepositoryImpl) DeleteProject(projectCode string) error {
	pjcd, err := repository.FindByIdProject(projectCode)
	if err != nil {
		return err
	}
	response := repository.DB.Delete(&domain.Project{}, pjcd)
	if response.Error != nil {
		return response.Error
	}
	return nil
}
