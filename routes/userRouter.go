package routes

import (
	controller "github.com/Muhammad-Hasan5/jwt-auth-GO/controllers"
	"github.com/Muhammad-Hasan5/jwt-auth-GO/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
