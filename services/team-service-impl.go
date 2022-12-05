package services

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/helper"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
	"github.com/go-playground/validator/v10"
)

type TeamServiceImpl struct {
	TeamRepository repositories.TeamRepository
	Validate       *validator.Validate
}

func NewServiceTeam(teamRepository repositories.TeamRepository, validate *validator.Validate) TeamService {
	return &TeamServiceImpl{
		TeamRepository: teamRepository,
		Validate:       validate,
	}
}

func (service *TeamServiceImpl) FindByIdTeam(teamCode string) (*response.TeamResponse, error) {
	team, err := service.TeamRepository.FindByIdTeam(teamCode)
	if err != nil {
		return nil, err
	}
	return helper.ToTeamResponse(team), nil
}

func (service *TeamServiceImpl) FindAllTeams() ([]*response.TeamResponse, error) {
	teams, err := service.TeamRepository.FindAllTeams()
	if err != nil {
		return nil, err
	}
	return helper.ToTeamResponses(teams), nil
}

func (service *TeamServiceImpl) CreateTeam(request *request.TeamRequest) (*response.TeamResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, err
	}
	team, err := service.TeamRepository.CreateTeam(&domain.Team{
		TeamCode:     request.TeamCode,
		TeamName:     request.TeamName,
		Leader:       request.Leader,
		ProjectBased: request.ProjectBased,
		Level:        request.Level,
		Status:       request.Status,
	})
	if err != nil {
		return nil, err
	}
	return helper.ToTeamResponse(team), nil
}

func (service *TeamServiceImpl) UpdateTeam(teamCode string, request *request.TeamRequest) (*response.TeamResponse, error) {
	_, err := service.TeamRepository.FindByIdTeam(teamCode)
	if err != nil {
		return nil, err
	}

	team, err := service.TeamRepository.UpdateTeam(teamCode, &domain.Team{
		TeamCode:     request.TeamCode,
		TeamName:     request.TeamName,
		Leader:       request.Leader,
		ProjectBased: request.ProjectBased,
		Level:        request.Level,
		Status:       request.Status,
	})
	if err != nil {
		return nil, err
	}
	return helper.ToTeamResponse(team), nil
}

func (service *TeamServiceImpl) DeleteTeam(teamCode string) error {
	err := service.TeamRepository.DeleteTeam(teamCode)
	if err != nil {
		return err
	}
	return nil
}
