package controllers

import (
	"net/http"

	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	teamcode := c.Param("teamcode")
	team, err := controller.TeamService.FindByIdTeam(teamcode)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "team code not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    team,
	})
}
func (controller *TeamControllerImpl) FindAllTeams(c *gin.Context) {
	teams, err := controller.TeamService.FindAllTeams()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": teams,
	})
}
