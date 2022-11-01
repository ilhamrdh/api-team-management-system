package helper

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
)

func ToTeamResponse(team domain.Team) web.TeamResponse {
	return web.TeamResponse{
		TeamCode:     team.TeamCode,
		TeamName:     team.TeamName,
		Leader:       team.Leader,
		ProjectBased: team.ProjectBased,
		Level:        team.Level,
		Status:       team.Status,
		CreatedAt:    team.CreatedAt,
		UpdatedAt:    team.UpdatedAt,
		DeletedAt:    team.DeletedAt,
	}
}

func ToTeamResponses(teams []domain.Team) []web.TeamResponse {
	var teamResponses []web.TeamResponse
	for _, team := range teams {
		teamResponses = append(teamResponses, ToTeamResponse(team))
	}
	return teamResponses
}

func ToProjectResponse(project domain.Project) web.ProjectResponse {
	return web.ProjectResponse{
		ProjectCode: project.ProjectCode,
		ProjectName: project.ProjectName,
		Deadline:    project.Deadline,
		Status:      project.Status,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		DeletedAt:   project.DeletedAt,
	}
}

func ToProjectResponses(projects []domain.Project) []web.ProjectResponse {
	var projectResponses []web.ProjectResponse
	for _, project := range projects {
		projectResponses = append(projectResponses, ToProjectResponse(project))
	}
	return projectResponses
}
