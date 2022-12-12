package controllers

import (
	"net/http"

	"github.com/IlhamRamadhan-IR/api-team-management-system/exceptions"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	ProjectService services.ProjectService
}

func NewControllerProject(r *gin.RouterGroup, projectService services.ProjectService) {
	projectController := ProjectController{
		ProjectService: projectService,
	}

	r.GET("/projects", projectController.FindAllProject)
	r.GET("/:project_code", projectController.FindByIdProject)
	r.POST("/", projectController.CreateProject)
	r.PUT("/:project_code", projectController.UpdateProject)
	r.DELETE("/:project_code", projectController.DeleteProject)
}

func (controller *ProjectController) FindByIdProject(c *gin.Context) {
	projectCode := c.Param("project_code")
	project, err := controller.ProjectService.FindByIdProject(projectCode)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}
	web.ResponseSuccess(c, project, "Get Data Successfully", http.StatusOK)
}

func (controller *ProjectController) FindAllProject(c *gin.Context) {
	projects, err := controller.ProjectService.FindAllProjects()
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}
	web.ResponseSuccess(c, projects, "Get Data Successfully", http.StatusOK)
}

func (controller *ProjectController) CreateProject(c *gin.Context) {
	var req request.ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		exceptions.EntityException(c, err.Error())
		return
	}

	project, err := controller.ProjectService.CreateProject(&req)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	web.ResponseSuccess(c, project, "Create Data Successfully", http.StatusOK)
}
func (controller *ProjectController) UpdateProject(c *gin.Context) {
	var req request.ProjectRequestUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		exceptions.EntityException(c, err.Error())
		return
	}
	projectCode := c.Param("project_code")
	project, err := controller.ProjectService.UpdateProject(projectCode, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	web.ResponseSuccess(c, project, "Update Data Successfully", http.StatusOK)
}

func (controller *ProjectController) DeleteProject(c *gin.Context) {
	projectCode := c.Param("project_code")
	err := controller.ProjectService.DeleteProject(projectCode)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}
	web.ResponseSuccess(c, projectCode, "Delete Data Successfully", http.StatusOK)
}
