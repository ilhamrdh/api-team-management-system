package controllers

import "github.com/gin-gonic/gin"

type TeamController interface {
	FindByIdTeam(c *gin.Context)
	FindAllTeams(c *gin.Context)
}
