package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/IlhamRamadhan-IR/api-team-management-system/exceptions"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web"
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/IlhamRamadhan-IR/api-team-management-system/securities"
	"github.com/IlhamRamadhan-IR/api-team-management-system/services"
	"github.com/IlhamRamadhan-IR/api-team-management-system/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthService
	UserService services.UserService
}

func NewControllerAuth(r *gin.RouterGroup, authService services.AuthService, userService services.UserService) {
	authController := AuthController{
		AuthService: authService,
		UserService: userService,
	}
	r.POST("/register", authController.DoRegister)
	// r.POST("/login", authController.DoLogin)
}

func (controller *AuthController) DoRegister(c *gin.Context) {
	var register request.UserRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		exceptions.EntityException(c, err.Error())
		return
	}

	check := utils.ValidateForm(&register)
	if check != "" {
		exceptions.BadRequestException(c, check)
		return
	}

	// var user *domain.User
	findUser, _ := controller.UserService.FindByUsername(register.Username)
	if findUser.Username != "" {
		exceptions.NotFoundException(c, errors.New("username already exists").Error())
		return
	}

	hash, err := securities.HashPassword(register.Password)
	if err != nil {
		exceptions.AppException(c, err.Error())
	}

	register.Username = strings.ToLower(strings.ReplaceAll(register.Username, " ", ""))
	register.Password = strings.ReplaceAll(register.Password, " ", "")
	register.Password = hash

	user, err := controller.AuthService.DoRegister(&register)
	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}
	web.ResponseSuccess(c, user, "Register Successfully", http.StatusOK)
}
