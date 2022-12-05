package repositories

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"gorm.io/gorm"
)

type TeamRepositoryImpl struct {
	DB *gorm.DB
}

func NewRepositoryTeam(db *gorm.DB) TeamRepository {
	return &TeamRepositoryImpl{
		DB: db,
	}
}

func (repository *TeamRepositoryImpl) FindByIdTeam(teamCode string) (*domain.Team, error) {
	var team *domain.Team
	err := repository.DB.Where("team_code = ?", teamCode).First(&team).Error
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (repository *TeamRepositoryImpl) FindAllTeams() ([]*domain.Team, error) {
	var teams []*domain.Team
	err := repository.DB.Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (repository *TeamRepositoryImpl) CreateTeam(team *domain.Team) (*domain.Team, error) {
	err := repository.DB.Create(&team).Error
	if err != nil {
		return nil, err
	}
	return team, nil
}
func (repository *TeamRepositoryImpl) UpdateTeam(teamCode string, team *domain.Team) (*domain.Team, error) {
	err := repository.DB.Model(&domain.Team{TeamCode: teamCode}).Updates(team).Find(team).Error
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (repository *TeamRepositoryImpl) DeleteTeam(teamCode string) error {
	tmcd, err := repository.FindByIdTeam(teamCode)
	if err != nil {
		return err
	}
	// response := repository.DB.Where("team_code = ?", teamCode).Delete(&domain.Team{})
	response := repository.DB.Delete(&domain.Team{}, tmcd)
	if response.Error != nil {
		return response.Error
	}
	return nil
}
