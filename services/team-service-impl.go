package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/helper"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
)

type TeamServiceImpl struct {
	TeamRepository repositories.TeamRepository
}

func NewServiceTeam(teamRepository repositories.TeamRepository) TeamService {
	return &TeamServiceImpl{
		TeamRepository: teamRepository,
	}
}

func (service *TeamServiceImpl) FindByIdTeam(teamcode string) (web.TeamResponse, error) {
	team, err := service.TeamRepository.FindByIdTeam(teamcode)
	return helper.ToTeamResponse(team), err
}
func (service *TeamServiceImpl) FindAllTeams() ([]web.TeamResponse, error) {
	teams, err := service.TeamRepository.FindAllTeams()
	return helper.ToTeamResponses(teams), err
}
