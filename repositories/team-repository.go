package repositories

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
)

type TeamRepository interface {
	FindByIdTeam(teamcode string) (domain.Team, error)
	FindAllTeams() ([]domain.Team, error)
}
