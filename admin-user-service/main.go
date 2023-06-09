package main

import (
	routesV1 "admin-user-service/api/v1"
	config "admin-user-service/pkg/common"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag"
)

// @title User Account
// @version 2
// @description Apis for login, signup
// @host localhost:10011
// @BasePath /user/v1
func main() {

	r := gin.Default()
	routesV1.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Start the server
	err := r.Run(config.Config.AdminUserServiceIP + ":" + config.Config.AdminUserServicePort)
	if err != nil {

	}
}
