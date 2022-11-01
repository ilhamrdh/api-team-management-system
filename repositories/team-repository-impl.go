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

func (repository *TeamRepositoryImpl) FindByIdTeam(teamcode string) (domain.Team, error) {
	var team domain.Team
	err := repository.DB.Where("team_code = ?", teamcode).First(&team).Error

	return team, err
}
func (repository *TeamRepositoryImpl) FindAllTeams() ([]domain.Team, error) {
	var teams []domain.Team
	err := repository.DB.Find(&teams).Error

	return teams, err
}
