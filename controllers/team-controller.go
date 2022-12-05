package controllers

import "github.com/gin-gonic/gin"

type TeamController interface {
	CreateTeam(c *gin.Context)
	UpdateTeam(c *gin.Context)
	DeleteTeam(c *gin.Context)
	FindByIdTeam(c *gin.Context)
	FindAllTeams(c *gin.Context)
}
