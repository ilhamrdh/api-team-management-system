package repositories

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
)

type TeamRepository interface {
	CreateTeam(team *domain.Team) (*domain.Team, error)
	UpdateTeam(teamCode string, team *domain.Team) (*domain.Team, error)
	DeleteTeam(teamCode string) error
	FindByIdTeam(teamcode string) (*domain.Team, error)
	FindAllTeams() ([]*domain.Team, error)
}
