package controller

import (
	"portfolio/dtos"
	"portfolio/models"
	"portfolio/services"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jWtService   services.JWTService
}

func LoginHandler(loginService services.LoginService,
	jWtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dtos.LoginCredentials
	var model models.Users
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		user_info, _ := model.JWTInfo(credential.Email, credential.Password)
		return controller.jWtService.GenerateToken(credential.Email, true, user_info.Id, user_info.IsAdmin)

	}
	return ""
}
