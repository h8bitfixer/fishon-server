package v1

import (
	"admin-user-service/internal/handler"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the API routes for version 1
func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/admin/v1")
	{
		v1.POST("/login", handler.LoginHandler)
	}
}
