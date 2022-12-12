package router

import (
	"fmt"

	"github.com/IlhamRamadhan-IR/api-team-management-system/config"
	"github.com/IlhamRamadhan-IR/api-team-management-system/controllers"
	"github.com/IlhamRamadhan-IR/api-team-management-system/middlewares"
	"github.com/IlhamRamadhan-IR/api-team-management-system/repositories"
	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) {
	r := gin.New()

	validate := validator.New()

	teamRepository := repositories.NewRepositoryTeam(db)
	teamService := services.NewServiceTeam(teamRepository, validate)

	projectRepository := repositories.NewRepositoryProject(db)
	projectService := services.NewServiceProject(projectRepository, validate)

	r.Use(middlewares.SetupCorsMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	controllers.NewControllerTeam(api.Group("/team"), teamService)
	controllers.NewControllerProject(api.Group("/project"), projectService)
	err := r.Run(fmt.Sprintf(":%v", config.Config.Server.Port))
	if err != nil {
		log.Fatal(err)
		return
	}
}
