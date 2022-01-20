package handler

import (
	"net/http"

	"portfolio/controller"
	"portfolio/services"

	"github.com/gin-gonic/gin"
)

func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginService services.LoginService = services.StaticLoginService()
		var jwtService services.JWTService = services.JWTAuthService()
		var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "username and password did not registered",
			})
		}

	}
}
