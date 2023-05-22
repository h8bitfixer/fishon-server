package v1

import (
	"github.com/gin-gonic/gin"
	"user-service/internal/handler"
)

// SetupRoutes sets up the API routes for version 1
func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/user/v1")
	{
		v1.POST("/signup", handler.SignupHandler)
		v1.GET("/login", handler.LoginHandler)
		v1.POST("/create-account", handler.CreateAccountHandler)
	}
}
