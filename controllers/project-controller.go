package controllers

import "github.com/gin-gonic/gin"

type ProjectController interface {
	FindByIdProject(c *gin.Context)
	FindAllProjects(c *gin.Context)
}
