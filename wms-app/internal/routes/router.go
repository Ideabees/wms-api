package routes

import (
	"wms-app/internal/controllers"
	"wms-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)

	r.POST("/login", controllers.Login)

	// Protected group
	protected := r.Group("/v1/api")
	protected.Use(utils.JWTMiddleware())
	{
		protected.GET("/permissions", controllers.GetPermissions)
		protected.POST("/logout", controllers.Logout)
	}
	return r
}
