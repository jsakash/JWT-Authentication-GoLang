package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jsakash/jwt-check/controllers"
	"github.com/jsakash/jwt-check/middleware"
)

func AuthRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/signup", controllers.Signup)
	incomingRoutes.POST("/login", controllers.Login)
	incomingRoutes.GET("/validate", middleware.RequireAuth, controllers.Validate)

}
