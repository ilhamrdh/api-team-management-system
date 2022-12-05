package controllers

import (
	"net/http"

	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
)

type TeamControllerImpl struct {
	TeamService services.TeamService
}

func NewControllerTeam(teamService services.TeamService) TeamController {
	return &TeamControllerImpl{
		TeamService: teamService,
	}
}

func (controller *TeamControllerImpl) FindByIdTeam(c *gin.Context) {
	teamCode := c.Param("team_code")

	team, err := controller.TeamService.FindByIdTeam(teamCode)
	if err != nil {
		errorRespose := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": errorRespose})
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   team,
	}

	c.JSON(http.StatusOK, gin.H{"data": webResponse})
}

func (controller *TeamControllerImpl) FindAllTeams(c *gin.Context) {
	teams, err := controller.TeamService.FindAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   teams,
	}

	c.JSON(http.StatusOK, gin.H{"data": webResponse})
}

func (controller *TeamControllerImpl) CreateTeam(c *gin.Context) {
	var req request.TeamRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	team, err := controller.TeamService.CreateTeam(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   team,
	}

	c.JSON(http.StatusOK, gin.H{"data": webResponse})
}

func (controller *TeamControllerImpl) UpdateTeam(c *gin.Context) {
	var req request.TeamRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	teamCode := c.Param("team_code")
	team, err := controller.TeamService.UpdateTeam(teamCode, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   team,
	}

	c.JSON(http.StatusOK, gin.H{"data": webResponse})
}

func (controller *TeamControllerImpl) DeleteTeam(c *gin.Context) {
	teamCode := c.Param("team_code")
	err := controller.TeamService.DeleteTeam(teamCode)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success deleted"})
}
