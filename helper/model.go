package helper

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/response"
)

func ToTeamResponse(team *domain.Team) *response.TeamResponse {
	return &response.TeamResponse{
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

func ToTeamResponses(teams []*domain.Team) []*response.TeamResponse {
	var teamResponses []*response.TeamResponse
	for _, team := range teams {
		teamResponses = append(teamResponses, ToTeamResponse(team))
	}
	return teamResponses
}

func ToProjectResponse(project *domain.Project) *response.ProjectResponse {
	return &response.ProjectResponse{
		ProjectCode: project.ProjectCode,
		ProjectName: project.ProjectName,
		Deadline:    project.Deadline,
		Status:      project.Status,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		DeletedAt:   project.DeletedAt,
	}
}

func ToProjectResponses(projects []*domain.Project) []*response.ProjectResponse {
	var projectResponses []*response.ProjectResponse
	for _, project := range projects {
		projectResponses = append(projectResponses, ToProjectResponse(project))
	}
	return projectResponses
}

func ToUserResponse(user *domain.User) *response.UserResponse {
	return &response.UserResponse{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func ToUserResponses(users []*domain.User) []*response.UserResponse {
	var userResponses []*response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToLoginResponse(user *domain.User) *response.LoginResponse {
	return &response.LoginResponse{
		Username: user.Username,
		Password: user.Password,
	}
}
func ToLoginResponses(users []*domain.User) []*response.LoginResponse {
	var userResponses []*response.LoginResponse
	for _, user := range users {
		userResponses = append(userResponses, ToLoginResponse(user))
	}
	return userResponses
}
