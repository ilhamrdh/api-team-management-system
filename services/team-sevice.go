package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
)

type TeamService interface {
	CreateTeam(request *request.TeamRequest) (*response.TeamResponse, error)
	UpdateTeam(teamCode string, request *request.TeamRequest) (*response.TeamResponse, error)
	DeleteTeam(teamCode string) error
	FindByIdTeam(teamCode string) (*response.TeamResponse, error)
	FindAllTeams() ([]*response.TeamResponse, error)
}
