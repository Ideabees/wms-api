package routes

import (
	"wms-app/internal/controllers"
	"wms-app/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your client's origin(s)
    config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.AllowCredentials = true
    r.Use(cors.New(config))

	r.POST("/api/register", controllers.Register)

	r.POST("/api/login", controllers.Login)

	// Protected group
	protected := r.Group("/v1/api")
	protected.Use(utils.JWTMiddleware())
	{
		protected.GET("/permissions", controllers.GetPermissions)
		protected.POST("/customer", controllers.CreateCustomer)
		protected.GET("/customers", controllers.GetCustomers)
		protected.DELETE("/customers", controllers.DeleteCustomers)
		protected.POST("/create_bulk_customers", controllers.CreateBulkCustomers)
		protected.POST("/logout", controllers.Logout)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
