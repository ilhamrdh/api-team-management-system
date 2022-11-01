package controllers

import (
	"net/http"

	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProjectControllerImpl struct {
	ProjectService services.ProjectService
}

func NewControllerProject(projectService services.ProjectService) ProjectController {
	return &ProjectControllerImpl{
		ProjectService: projectService,
	}
}

func (controller *ProjectControllerImpl) FindByIdProject(c *gin.Context) {
	projectcode := c.Param("projectcode")
	project, err := controller.ProjectService.FindByIdProject(projectcode)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "project code not found",
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
		"data":    project,
	})
}
func (controller *ProjectControllerImpl) FindAllProjects(c *gin.Context) {
	projects, err := controller.ProjectService.FindAllProjects()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messsage": "OK",
		"data":     projects,
	})
}
