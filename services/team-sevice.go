package services

import "github.com/IlhamRamadhan-IR/api-team-management-system/models/web"

type TeamService interface {
	FindByIdTeam(teamcode string) (web.TeamResponse, error)
	FindAllTeams() ([]web.TeamResponse, error)
}
