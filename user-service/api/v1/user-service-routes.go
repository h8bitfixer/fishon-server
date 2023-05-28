package v1

import (
	"github.com/gin-gonic/gin"
	"user-service/internal/handler"
	middleware "user-service/internal/service"
)

// SetupRoutes sets up the API routes for version 1
func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/user/v1")
	{
		v1.POST("/get_otp", handler.GetOTPHandler)
		v1.POST("/verify_otp", handler.VerifyOTPHandler)

		v1.GET("/login", middleware.JWTAuth(), handler.LoginHandler)
		v1.POST("/create-account", middleware.JWTAuth(), handler.CreateAccountHandler)
	}
}
