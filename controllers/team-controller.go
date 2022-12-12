package controllers

import (
	"net/http"

	"github.com/IlhamRamadhan-IR/api-team-management-system/exceptions"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	TeamService services.TeamService
}

func NewControllerTeam(r *gin.RouterGroup, teamService services.TeamService) {
	teamController := TeamController{
		TeamService: teamService,
	}
	r.GET("/teams", teamController.FindAllTeams)
	r.GET("/:team_code", teamController.FindByIdTeam)
	r.POST("/", teamController.CreateTeam)
	r.PUT("/:team_code", teamController.UpdateTeam)
	r.DELETE("/:team_code", teamController.DeleteTeam)
}

func (controller *TeamController) FindByIdTeam(c *gin.Context) {
	teamCode := c.Param("team_code")

	team, err := controller.TeamService.FindByIdTeam(teamCode)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}
	web.ResponseSuccess(c, team, "Get Data Successfully", http.StatusOK)
}

func (controller *TeamController) FindAllTeams(c *gin.Context) {
	teams, err := controller.TeamService.FindAllTeams()
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}
	web.ResponseSuccess(c, teams, "Get Data Successfully", http.StatusOK)
}

func (controller *TeamController) CreateTeam(c *gin.Context) {
	var req request.TeamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		exceptions.EntityException(c, err.Error())
		return
	}

	team, err := controller.TeamService.CreateTeam(&req)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	web.ResponseSuccess(c, team, "Create Data Successfully", http.StatusOK)
}

func (controller *TeamController) UpdateTeam(c *gin.Context) {
	var req request.TeamRequestUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		exceptions.EntityException(c, err.Error())
		return
	}

	teamCode := c.Param("team_code")
	team, err := controller.TeamService.UpdateTeam(teamCode, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	web.ResponseSuccess(c, team, "Update Data Successfully", http.StatusOK)
}

func (controller *TeamController) DeleteTeam(c *gin.Context) {
	teamCode := c.Param("team_code")
	err := controller.TeamService.DeleteTeam(teamCode)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	web.ResponseSuccess(c, teamCode, "Delete Data Successfully", http.StatusOK)
}
