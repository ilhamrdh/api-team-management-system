package controllers

import (
	"net/http"
	"strconv"

	"github.com/IlhamRamadhan-IR/api-team-management-system/exceptions"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewControllerUser(r *gin.RouterGroup, userService services.UserService) {
	userController := UserController{
		UserService: userService,
	}

	r.GET("/:user_id", userController.FindById)
	r.GET("/username/:username", userController.FindUsername)
}

func (controller *UserController) FindUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := controller.UserService.FindByUsername(username)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}
	web.ResponseSuccess(c, user, "Get Data By Username Successfully", http.StatusOK)
}
func (controller *UserController) FindById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("user_id"))
	user, err := controller.UserService.FindById(userId)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}
	web.ResponseSuccess(c, user, "Get Data By id Successfully", http.StatusOK)
}
