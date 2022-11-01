package main

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/app"
	"github.com/IlhamRamadhan-IR/api-team-management-system/controllers"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := app.InitDatabase()

	teamRepository := repositories.NewRepositoryTeam(db)
	teamService := services.NewServiceTeam(teamRepository)
	teamController := controllers.NewControllerTeam(teamService)

	projectRepository := repositories.NewRepositoryProject(db)
	projectService := services.NewServiceProject(projectRepository)
	projectController := controllers.NewControllerProject(projectService)

	r := gin.Default()
	api := r.Group("/api")
	{
		team := api.Group("/team")
		{
			team.GET("/getteams", teamController.FindAllTeams)
			team.GET("/getteam/:teamcode", teamController.FindByIdTeam)
		}
		project := api.Group("/project")
		{
			project.GET("/getprojects", projectController.FindAllProjects)
			project.GET("/getproject/:projectcode", projectController.FindByIdProject)
		}
	}
	r.Run()
}
