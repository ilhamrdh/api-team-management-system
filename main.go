package main

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/app"
	"github.com/IlhamRamadhan-IR/api-team-management-system/controllers"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.InitDatabase()
	validate := validator.New()
	teamRepository := repositories.NewRepositoryTeam(db)
	teamService := services.NewServiceTeam(teamRepository, validate)
	teamController := controllers.NewControllerTeam(teamService)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api")
	{
		team := api.Group("/team")
		{
			team.GET("/teams", teamController.FindAllTeams)
			team.GET("/:team_code", teamController.FindByIdTeam)
			team.POST("/", teamController.CreateTeam)
			team.PUT("/:team_code", teamController.UpdateTeam)
			team.DELETE("/:team_code", teamController.DeleteTeam)
		}
	}
	r.Run()
}
