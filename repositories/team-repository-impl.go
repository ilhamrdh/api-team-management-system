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
	err := repository.DB.Model(&domain.Team{}).Where("team_code = ?", teamCode).First(&team).Error
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (repository *TeamRepositoryImpl) FindAllTeams() ([]*domain.Team, error) {
	var teams []*domain.Team
	err := repository.DB.Model(&domain.Team{}).Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (repository *TeamRepositoryImpl) CreateTeam(team *domain.Team) (*domain.Team, error) {
	tx := repository.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}
	err := tx.Create(&team).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return team, nil
}
func (repository *TeamRepositoryImpl) UpdateTeam(teamCode string, team *domain.Team) (*domain.Team, error) {
	err := repository.DB.Exec("UPDATE team SET team_name=?,leader=?,project_based=?,level=?,status=? WHERE team_code=?",
		team.TeamName,
		team.Leader,
		team.ProjectBased,
		team.Level,
		team.Status,
		teamCode,
	).Error
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (repository *TeamRepositoryImpl) DeleteTeam(teamCode string) error {
	code, err := repository.FindByIdTeam(teamCode)
	if err != nil {
		return err
	}
	// response := repository.DB.Where("team_code = ?", teamCode).Delete(&domain.Team{})
	response := repository.DB.Model(&domain.Team{}).Delete(&domain.Team{}, code)
	if response.Error != nil {
		return response.Error
	}
	return nil
}
